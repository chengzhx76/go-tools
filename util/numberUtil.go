package util

import (
	"fmt"
	. "github.com/chengzhx76/go-tools/consts"
	"math"
	"strings"
)

// 将浮点数四舍五入为 string，downDouble 为缩小的倍数
func Round(num int64, downDouble int) string {
	//return StringToFloat64(fmt.Sprintf("%.2f", float64(num)/float64(downDouble)))
	return fmt.Sprintf("%.2f", float64(num)/float64(downDouble))
}

/*
	函数功能：将浮点数 val 四舍五入到指定的小数位数 precision。

	参数说明：

	val：需要四舍五入的浮点数
	precision：保留的小数位数

	RoundFloat(3.14159, 2) → 3.14
	RoundFloat(3.14159, 4) → 3.1416
	RoundFloat(3.145, 2) → 3.15（四舍五入）
	RoundFloat(3.144, 2) → 3.14
	RoundFloat(123.456, 0) → 123（保留0位小数，即取整）

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

/*
函数功能：将浮点数 val 四舍五入后转换为 int64 类型的整数。
RoundInt64(3.14) → 3
RoundInt64(3.5) → 4（四舍五入）
RoundInt64(3.9) → 4
RoundInt64(-2.7) → -3
RoundInt64(-2.3) → -2
*/
func RoundInt64(val float64) int64 {
	return int64(RoundFloat(val, 0))
}

// https://blog.csdn.net/qq_42410605/article/details/125339144

// 向上取整 1.11->2, 1.99->2, -1.11->1
func RoundUp(val float64) int64 {
	return int64(math.Ceil(val))
}

// 向下取整 1.11->1, 1.99->1, -1.11->2
func RoundDown(val float64) int64 {
	return int64(math.Floor(val))
}

// NumberFormatCn(20001) // 2万
// NumberFormatCn(20100) // 2万1千
// NumberFormatCn(20199) // 2万1千
// NumberFormatCn(21900) // 2万2千
func NumberFormatCn(num int64) string {
	numStr := Int64ToString(num)
	if len(numStr) == 4 {
		/*maxNum := StringToInt64(strings.Replace(numStr, SubString(numStr, -3, 0), SYMBOL_EMPTY, 1))
		decimalPart := SubString(numStr, -3, 2)
		decimal := RoundUp(StringToFloat64(decimalPart) / float64(10))
		if decimal == 0 {
			return fmt.Sprintf("%d千", maxNum)
		} else if decimal == 10 {
			maxNum += 1
			return fmt.Sprintf("%d千", maxNum)
		}
		return fmt.Sprintf("%d千%d百", maxNum, decimal)*/

	} else if len(numStr) >= 5 {
		maxNum := StringToInt64(strings.Replace(numStr, SubString(numStr, -4, 0), SYMBOL_EMPTY, 1))
		decimalPart := SubString(numStr, -4, 2)
		decimal := RoundUp(StringToFloat64(decimalPart) / float64(10))
		if decimal == 0 {
			return fmt.Sprintf("%d万", maxNum)
		} else if decimal == 10 {
			maxNum += 1
			return fmt.Sprintf("%d万", maxNum)
		}
		return fmt.Sprintf("%d.%d万", maxNum, decimal)

	}
	return Int64ToString(num)
}

/*
函数功能：将 numToRound 向上取整到最接近的 multiple 的倍数。

参数说明：
numToRound：需要取整的数字
multiple：倍数基准

RoundUpMultiple(10, 3) → 12（10除以3余1，需要加2）
RoundUpMultiple(15, 5) → 15（15除以5余0，已经是5的倍数）
RoundUpMultiple(17, 4) → 20（17除以4余1，需要加3）
RoundUpMultiple(1, 10) → 10（1除以10余1，需要加9）
*/
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
