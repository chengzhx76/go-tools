package backoff

import (
	"log"
	"math/rand"
	"time"
)

// Backoff 定义了指数退避策略的结构体
type Backoff struct {
	maxRetries int
	baseDelay  time.Duration
	maxDelay   time.Duration
	RNG        *rand.Rand
}

// NewBackoff 创建一个新的 Backoff 实例
// maxRetries: 最大重试次数
// baseDelay: 初始延迟时间
// maxDelay: 最大延迟时间
func NewBackoff(maxRetries int, baseDelay, maxDelay time.Duration) *Backoff {
	return &Backoff{
		maxRetries: maxRetries,
		baseDelay:  baseDelay,
		maxDelay:   maxDelay,
		RNG:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Perform 尝试执行指定的操作，使用指数退避策略进行重试
// operationCb: 需要执行的操作回调函数
// arg: 传递给操作回调函数的参数
func (b *Backoff) Perform(operationCb func(any) error, arg any) {
	go func() {
		retries := 0

		for retries < b.maxRetries {
			err := operationCb(arg)
			if err == nil {
				return
			}

			retries++
			delay := b.baseDelay * (1 << retries)
			if delay > b.maxDelay {
				delay = b.maxDelay
			}

			// 添加随机抖动
			jitter := time.Duration(b.RNG.Int63n(int64(delay) / 2))
			time.Sleep(delay + jitter)
		}
		log.Println("Operation failed after maximum retries")
	}()
}

/*
func main() {
    rand.Seed(time.Now().UnixNano())

    // 创建一个 Backoff 实例
    b := backoff.NewBackoff(5, 1*time.Second, 32*time.Second)

    // 定义需要重试的操作
    operation := func() error {
        // 这里模拟一个可能会失败的操作，实际使用时替换成真实操作
        if rand.Float32() < 0.8 {
            return errors.New("simulated operation failure")
        }
        return nil
    }

    // 使用 Backoff 执行操作
    err := b.Perform(operation)
    if err != nil {
        fmt.Println("Operation failed:", err)
    } else {
        fmt.Println("Operation succeeded")
    }
}
*/
