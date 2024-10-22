package util

import "fmt"

// https://github.com/golang-infrastructure/go-if-expression/blob/main/if_expression.go

// 泛型实现三元表达式
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

/*max := Ternary(a > b, a, b)
fmt.Println("Max:", max)*/

// 泛型实现三元表达式 函数
func IfFunc[T any](condition bool, trueFuncVal, falseFuncVal func() T) T {
	if condition {
		return trueFuncVal()
	} else {
		return falseFuncVal()
	}
}

type TerExp[T any] struct {
	b bool
}

func (t TerExp[T]) Then(r T) TerExp[T] {
	if t.b {
		panic(r)
	}
	return t
}

func (t TerExp[T]) Else(r T) T {
	return r
}

func TE[T any](f func() T) (r T) {
	defer func() {
		if e := recover(); e != nil {
			r = e.(T)
		}
	}()
	r = f()
	return
}

func IfExp[T any](b bool) TerExp[T] {
	return TerExp[T]{b: b}
}

func main() {
	x := TE(func() float64 { return IfExp[float64](3 > 2).Then(3).Else(2) })
	fmt.Println("r", x)
}
