package tests

import (
	"testing"
	"time"
)

func Test_nil(t *testing.T) {
	//tm := util.ParseLocalTime(time.RFC3339, "0001-01-01T00:00:00Z")
	LOCAL_TIME, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation(time.RFC3339, "0001-01-01T00:00:00Z", LOCAL_TIME)
	t.Log(tm)
}
