package util

import (
	"fmt"
	"math"
)

// 将浮点数四舍五入为 string，downDouble 为缩小的倍数
func Round(num int64, downDouble int) string {
	//return StringToFloat64(fmt.Sprintf("%.2f", float64(num)/float64(downDouble)))
	return fmt.Sprintf("%.2f", float64(num)/float64(downDouble))
}

/*
	https://gosamples.dev/round-float/

	number := 12.3456789

	fmt.Println(roundFloat(number, 2))
	fmt.Println(roundFloat(number, 3))
	fmt.Println(roundFloat(number, 4))
	fmt.Println(roundFloat(number, 5))

	number = -12.3456789
	fmt.Println(roundFloat(number, 0))
	fmt.Println(roundFloat(number, 1))
	fmt.Println(roundFloat(number, 10))

Output:

	12.35
	12.346
	12.3457
	12.34568
	-12
	-12.3
	-12.3456789
*/
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
func RoundInt64(val float64) int64 {
	return int64(RoundFloat(val, 0))
}

// https://blog.csdn.net/qq_42410605/article/details/125339144

// 向上取整 1.11->2, 1.99->2, -1.11->1
func RoundUp(val float64) int64 {
	return int64(math.Ceil(val))
}

// 向上取整 1.11->1, 1.99->1, -1.11->2
func RoundDown(val float64) int64 {
	return int64(math.Floor(val))
}

func RoundUpMultiple(numToRound, multiple int64) int64 {
	if multiple == 0 {
		return numToRound
	}
	remainder := numToRound % multiple
	if remainder == 0 {
		return numToRound
	}
	return numToRound + multiple - remainder
}
