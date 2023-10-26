package util

import (
	"fmt"
)

func Round(num int64, downDouble int) string {
	//return StringToFloat64(fmt.Sprintf("%.2f", float64(num)/float64(downDouble)))
	return fmt.Sprintf("%.2f", float64(num)/float64(downDouble))
}
