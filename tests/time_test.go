package tests

import (
	"fmt"
	"github.com/chengzhx76/go-tools/consts"
	"github.com/chengzhx76/go-tools/util"
	"testing"
	"time"
)

func Test_nil(t *testing.T) {
	//tm := util.ParseLocalTime(time.RFC3339, "0001-01-01T00:00:00Z")
	LOCAL_TIME, _ := time.LoadLocation("Local")
	tm, _ := time.ParseInLocation(time.RFC3339, "0001-01-01T00:00:00Z", LOCAL_TIME)
	t.Log(tm)
	t.Log(tm.Format(consts.DATE_TIME_FORMAT))
}

func Test_Differ(t *testing.T) {
	start := util.ParseLocalTime(consts.DATE_TIME_FORMAT, "2024-11-10 22:59:59")
	end := util.ParseLocalTime(consts.DATE_TIME_FORMAT, "2024-11-10 23:59:59")

	s := util.DayDiffer(end, start)

	t.Logf("===> %v", s)

}

func Test_AddDay(t *testing.T) {

	s := util.AddDay(1, util.EndOfDay(time.Now()))
	num := util.DayDiffer(s, time.Now())

	t.Logf("===> %v %d", s, num)

}

func Test_Differ_s(t *testing.T) {
	start := util.ParseLocalTime(consts.DATE_TIME_FORMAT, "2024-11-12 01:59:59")
	end := util.ParseLocalTime(consts.DATE_TIME_FORMAT, "2024-11-14 23:59:59")

	seconds := util.EndOfDay(end).Sub(util.EndOfDay(start)).Seconds()
	if seconds == 0 {
		//return 1
	} else {

	}
	//seconds := end.Sub(start).Seconds()
	t.Logf("===> %v|%v|%v|%v|%v", seconds, 24*60*60, fmt.Sprintf("%.2f", seconds/float64(24*60*60)), util.RoundUp(seconds/(24*60*60)), seconds/(24*60*60))
}
