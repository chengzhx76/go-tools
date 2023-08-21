package util

import "regexp"

var MOBILE = regexp.MustCompile("(?:0|86|\\+86)?1[3-9]\\d{9}")

func IsMobile(str string) bool {
	return MOBILE.MatchString(str)
}
