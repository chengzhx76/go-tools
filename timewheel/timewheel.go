package timewheel

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// Timeout 哈希轮 / Calendar Queue 的实现，用于定时任务调度。
// 通过 time.Ticker 驱动定时滴答。
const (
	defaultTickInterval = time.Millisecond // 默认的时间粒度
	defaultNumBuckets   = 2048             // 默认桶数量

	cacheline    = 64
	bitsInUint64 = 64
)

const (
	// TimeoutWheel 的生命周期状态
	stopped int32 = iota
	stopping
	running
)

const (
	// 单个 Timeout 的状态
	timeoutInactive = iota
	timeoutExpired
	timeoutActive
)

var (
	// ErrSystemStopped 在用户尝试在系统停止后继续调度时返回
	ErrSystemStopped = errors.New("Timeout System is stopped")

	// 4*(NumCPU 向下取整到最接近的 2 的幂次方)
	defaultPoolSize = uint64(1 << uint(findMSB(uint64(runtime.NumCPU()))+2))
)

// Timeout 表示一个待触发的回调函数。
type Timeout struct {
	generation uint64
	timeout    *timeout
}

// Stop 停止已调度的超时，防止回调被执行。返回 true 表示取消成功。
func (t *Timeout) Stop() bool {
	if t.timeout == nil {
		return false
	}

	t.timeout.mtx.Lock()
	if t.timeout.generation != t.generation || t.timeout.state != timeoutActive {
		t.timeout.mtx.Unlock()
		return false
	}

	// 从桶中移除并放回空闲链表
	t.timeout.removeLocked()
	t.timeout.wheel.putTimeoutLocked(t.timeout)
	t.timeout.mtx.Unlock()
	return true
}

type timeout struct {
	mtx       *paddedMutex
	expireCb  func(any)
	expireArg any
	deadline  uint64

	// 前向指针以及前驱的 next 指针地址。通过这种方式实现 O(1) 删除。
	next  *timeout
	prev  **timeout
	state int32

	wheel      *TimeoutWheel
	generation uint64
}

type timeoutList struct {
	lastTick uint64
	head     *timeout
}

func (t *timeout) prependLocked(list *timeoutList) {
	if list.head != nil {
		t.prev = list.head.prev
		list.head.prev = &t.next
	} else {
		t.prev = &list.head
	}
	t.next = list.head
	list.head = t
}

func (t *timeout) removeLocked() {
	if t.next != nil {
		t.next.prev = t.prev
	}

	*t.prev = t.next
	t.next = nil
	t.prev = nil
}

// TimeoutWheel 是按桶划分的定时器集合（当前 tick 精度为 1ms）。
type TimeoutWheel struct {
	// ticks 为全局滴答计数器
	ticks uint64

	// buckets[i] 和 freelists[i] 的锁由 mtxPool[i&poolMask] 控制
	mtxPool      []paddedMutex
	bucketMask   uint64
	poolMask     uint64
	tickInterval time.Duration
	buckets      []timeoutList
	freelists    []timeoutList

	state     int32
	calloutCh chan timeoutList
	done      chan struct{}
}

// Option 用于配置 NewTimeoutWheel
type Option func(*opts)

type opts struct {
	tickInterval time.Duration
	size         uint64
	poolsize     uint64
}

// 加入填充以避免 mutex 发生伪共享
type paddedMutex struct {
	sync.Mutex
	_ [cacheline - unsafe.Sizeof(sync.Mutex{})]byte
}

// WithTickInterval 设置滴答间隔
func WithTickInterval(interval time.Duration) Option {
	return func(opts *opts) { opts.tickInterval = interval }
}

// WithBucketsExponent 设置桶数量的幂指数
func WithBucketsExponent(bucketExp uint) Option {
	return func(opts *opts) {
		opts.size = uint64(1 << bucketExp)
	}
}

// WithLocksExponent 设置锁池数量的幂指数，当大于桶数量时自动取桶数量
func WithLocksExponent(lockExp uint) Option {
	return func(opts *opts) {
		opts.poolsize = uint64(1 << lockExp)
	}
}

// NewTimeoutWheel 创建并启动一个 TimeoutWheel。
func NewTimeoutWheel(options ...Option) *TimeoutWheel {
	opts := &opts{
		tickInterval: defaultTickInterval,
		size:         defaultNumBuckets,
		poolsize:     defaultPoolSize,
	}

	for _, option := range options {
		option(opts)
	}

	poolsize := opts.poolsize
	if opts.size < opts.poolsize {
		poolsize = opts.size
	}

	t := &TimeoutWheel{
		mtxPool:      make([]paddedMutex, poolsize),
		freelists:    make([]timeoutList, poolsize),
		state:        stopped,
		poolMask:     poolsize - 1,
		buckets:      make([]timeoutList, opts.size),
		tickInterval: opts.tickInterval,
		bucketMask:   opts.size - 1,
		ticks:        0,
	}
	t.Start()
	return t
}

func (t *TimeoutWheel) getState() int32 {
	return atomic.LoadInt32(&t.state)
}

func (t *TimeoutWheel) updateState(state int32) {
	atomic.StoreInt32(&t.state, state)
}

// Start 启动一个已停止的 TimeoutWheel。重复调用会导致 panic。
func (t *TimeoutWheel) Start() {
	t.lockAllBuckets()
	defer t.unlockAllBuckets()

	for t.getState() != stopped {
		switch t.getState() {
		case stopping:
			// 如果正在停止，等待停止流程完成
			t.unlockAllBuckets()
			<-t.done
			t.lockAllBuckets()
		case running:
			panic("Tried to start a running TimeoutWheel")
		}
	}

	// 初始化运行状态和通道
	t.updateState(running)
	t.done = make(chan struct{})
	t.calloutCh = make(chan timeoutList)

	go t.doTick()
	go t.doExpired()
}

// Stop 停止滴答并清理剩余超时任务。
func (t *TimeoutWheel) Stop() {
	t.lockAllBuckets()

	if t.getState() == running {
		t.updateState(stopping)
		close(t.calloutCh)
		for i := range t.buckets {
			t.freeBucketLocked(t.buckets[i])
		}
	}

	// 解锁以便回调 goroutine 完成
	t.unlockAllBuckets()
	<-t.done
}

// Schedule 在持续时间 d 后调度执行回调函数。若 d 介于两个滴答之间，将归入后一个滴答。
func (t *TimeoutWheel) Schedule(d time.Duration, expireCb func(any), arg any) (Timeout, error) {
	dTicks := (d + t.tickInterval - 1) / t.tickInterval
	deadline := atomic.LoadUint64(&t.ticks) + uint64(dTicks)
	timeout := t.getTimeoutLocked(deadline)

	if t.getState() != running {
		// 若当前未运行，直接回收 timeout
		t.putTimeoutLocked(timeout)
		timeout.mtx.Unlock()
		return Timeout{}, ErrSystemStopped
	}

	bucket := &t.buckets[deadline&t.bucketMask]
	timeout.expireCb = expireCb
	timeout.expireArg = arg
	timeout.deadline = deadline
	timeout.state = timeoutActive
	out := Timeout{timeout: timeout, generation: timeout.generation}

	// 若该桶的 lastTick 已超过 deadline，则立即执行回调
	if bucket.lastTick >= deadline {
		t.putTimeoutLocked(timeout)
		timeout.mtx.Unlock()
		expireCb(arg)
		return out, nil
	}

	// 插入到目标桶的链表头
	timeout.prependLocked(bucket)
	timeout.mtx.Unlock()
	return out, nil
}

// doTick 负责处理滴答的 goroutine。
func (t *TimeoutWheel) doTick() {
	var expiredList timeoutList

	ticker := time.NewTicker(t.tickInterval)
	for range ticker.C {
		v := atomic.AddUint64(&t.ticks, 1)

		// 锁定对应桶，保证对同一桶的安全访问
		mtx := t.lockBucket(v)
		if t.getState() != running {
			mtx.Unlock()
			break
		}

		bucket := &t.buckets[v&t.bucketMask]
		timeout := bucket.head
		bucket.lastTick = v

		// 遍历该桶，找出已到期的 timeout
		for timeout != nil {
			next := timeout.next
			if timeout.deadline <= v {
				timeout.state = timeoutExpired
				timeout.removeLocked()
				timeout.prependLocked(&expiredList)
			}
			timeout = next
		}

		mtx.Unlock()
		if expiredList.head == nil {
			continue
		}

		// 将已到期的任务转移到 callout 通道等待回调处理
		select {
		case t.calloutCh <- expiredList:
			expiredList.head = nil
		default:
			// 如果通道已满，留待下次循环继续发送
		}
	}

	ticker.Stop()
}

func (t *TimeoutWheel) getTimeoutLocked(deadline uint64) *timeout {
	mtx := &t.mtxPool[deadline&t.poolMask]
	mtx.Lock()
	freelist := &t.freelists[deadline&t.poolMask]
	if freelist.head == nil {
		timeout := &timeout{mtx: mtx, wheel: t}
		return timeout
	}
	timeout := freelist.head
	timeout.removeLocked()
	return timeout
}

func (t *TimeoutWheel) putTimeoutLocked(timeout *timeout) {
	freelist := &t.freelists[timeout.deadline&t.poolMask]
	timeout.state = timeoutInactive
	timeout.generation++
	timeout.prependLocked(freelist)
}

func (t *TimeoutWheel) lockBucket(bucket uint64) *paddedMutex {
	mtx := &t.mtxPool[bucket&t.poolMask]
	mtx.Lock()
	return mtx
}

func (t *TimeoutWheel) lockAllBuckets() {
	for i := range t.mtxPool {
		t.mtxPool[i].Lock()
	}
}

func (t *TimeoutWheel) unlockAllBuckets() {
	for i := len(t.mtxPool) - 1; i >= 0; i-- {
		t.mtxPool[i].Unlock()
	}
}

func (t *TimeoutWheel) freeBucketLocked(head timeoutList) {
	timeout := head.head
	for timeout != nil {
		next := timeout.next
		timeout.removeLocked()
		t.putTimeoutLocked(timeout)
		timeout = next
	}
}

// doExpired 负责回调执行和资源回收
func (t *TimeoutWheel) doExpired() {
	for list := range t.calloutCh {
		timeout := list.head
		for timeout != nil {
			timeout.mtx.Lock()
			next := timeout.next
			expireCb := timeout.expireCb
			expireArg := timeout.expireArg
			t.putTimeoutLocked(timeout)
			timeout.mtx.Unlock()

			if expireCb != nil {
				// 执行用户回调
				expireCb(expireArg)
			}
			timeout = next
		}
	}

	// 确保结束时状态恢复为 stopped
	t.lockAllBuckets()
	t.updateState(stopped)
	t.unlockAllBuckets()
	close(t.done)
}

// findMSB 返回 value 的最高有效位位置（从 0 开始）
func findMSB(value uint64) int {
	for i := bitsInUint64 - 1; i >= 0; i-- {
		if value&(1<<uint(i)) != 0 {
			return int(i)
		}
	}
	return -1
}
