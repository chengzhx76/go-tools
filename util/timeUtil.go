package util

import (
	"fmt"
	. "github.com/chengzhx76/go-tools/consts"
	"log"
	"math"
	"strconv"
	"time"
)

func AddDay(offset int, dateTime time.Time) time.Time {
	return dateTime.Local().AddDate(0, 0, offset)
}

func AddWeek(offset int, dateTime time.Time) time.Time {
	return AddDay(7*offset, dateTime)
}

func AddMonth(offset int, dateTime time.Time) time.Time {
	addAfterTime := dateTime.Local().AddDate(0, offset, 0)
	if dateTime.Day() == addAfterTime.Day() {
		return addAfterTime
	}

	offsetAfterMonthFirstDate := GetFirstDateOfMonth(dateTime).Local().AddDate(0, offset, 0)
	offsetAfterMonthLastDate := GetLastDateOfMonth(offsetAfterMonthFirstDate)
	hour, min, sec := addAfterTime.Clock()
	return time.Date(offsetAfterMonthFirstDate.Year(), offsetAfterMonthFirstDate.Month(), offsetAfterMonthLastDate.Day(), hour, min, sec, dateTime.Nanosecond(), dateTime.Location())
}

func AddQuarter(offset int, dateTime time.Time) time.Time {
	return AddMonth(3*offset, dateTime)
}

func AddYear(offset int, dateTime time.Time) time.Time {
	addAfterTime := dateTime.Local().AddDate(offset, 0, 0)
	if dateTime.Day() == addAfterTime.Day() {
		return addAfterTime
	}
	offsetAfterMonthFirstDate := GetFirstDateOfMonth(dateTime).Local().AddDate(offset, 0, 0)
	offsetAfterMonthLastDate := GetLastDateOfMonth(offsetAfterMonthFirstDate)

	hour, min, sec := addAfterTime.Clock()
	return time.Date(offsetAfterMonthFirstDate.Year(), offsetAfterMonthFirstDate.Month(), offsetAfterMonthLastDate.Day(), hour, min, sec, dateTime.Nanosecond(), dateTime.Location())
}

func AddHour(offset int64, dateTime time.Time) time.Time {
	return dateTime.Local().Add(time.Hour * time.Duration(offset))
}

func AddMinute(offset int64, dateTime time.Time) time.Time {
	return dateTime.Local().Add(time.Minute * time.Duration(offset))
}

func AddSecond(offset int64, dateTime time.Time) time.Time {
	return dateTime.Local().Add(time.Second * time.Duration(offset))
}

// https://golang.org/pkg/time/#Time.AddDate
// 获取昨天
func Yesterday() time.Time {
	return AddDay(-1, time.Now())
}

// 获取今天
func Today() time.Time {
	return time.Now()
}

// 获取明天
func Tomorrow() time.Time {
	return AddDay(1, time.Now())
}

// 后天
func AfterTomorrow() time.Time {
	return AddDay(2, time.Now())
}

// 大后天
func BigAfterTomorrow() time.Time {
	return AddDay(3, time.Now())
}

// 获取某一天的0点时间
func GetZeroTime(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, dateTime.Location())
}

// 获取某年的第一天0点时间
func YearFirstDay(year string) time.Time {
	dateFull := fmt.Sprintf("%s0101", year)
	dateTime := ParseLocalTime(DATE_FORMAT_YYYYMMDD, dateFull)
	return GetZeroTime(dateTime)
}

// 获取某年的最后一天23:59点时间
func YearLastDay(year string) time.Time {
	dateFull := fmt.Sprintf("%s1231", year)
	dateTime := ParseLocalTime(DATE_FORMAT_YYYYMMDD, dateFull)
	return GetLastTime(dateTime)
}

// 获取某一天指定的时间
func GetSpecifyHourAndMinTime(dateTime time.Time, hour, min int) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), hour, min, 0, 0, dateTime.Location())
}

// 获取某一分钟的0秒时间
func GetMinStartTime(dateTime time.Time) time.Time {
	return GetSpecifyHourAndMinTime(dateTime, dateTime.Hour(), dateTime.Minute())
}

// 获取某一天的最后时间
func GetLastTime(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 23, 59, 59, 0, dateTime.Location())
}

// 获取当前日期的年份
func GetYear(date time.Time) int {
	return date.Year()
}

// 返回当前日期是星期几
func GetWeek(dateTime time.Time) int {
	week := int(dateTime.Weekday())
	if week == 0 {
		return 7 // 周日改为 7
	}
	return week
}

// 获取当前日期是几号
func GetDay(date time.Time) int {
	return date.Day()
}

// 获取从当前 date 算有之后的 cycle 个月有几天
func GetMonthDayNums(cycle int, date time.Time) int {
	days := 0
	for i := 0; i < cycle; i++ {
		date = AddMonth(1, date)
		days += GetLastDateOfMonth(date).Day()
	}
	return days
}

// dateTime 距离 weekNum（周几）还有几天
func FromWeekDays(weekNum int, dateTime time.Time) int {
	todayWeek := GetWeek(dateTime)
	offset := 0
	if weekNum >= todayWeek {
		offset = weekNum - todayWeek
	} else {
		offset = 7 - todayWeek + weekNum
	}
	return offset
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取num个月的日期数
func getNumMonthDays(num int, date time.Time) []string {
	differDays := DayDiffer(GetLastDateOfMonth(AddMonth(num, date)), date)
	var days []string
	var i int
	for i = 0; i < differDays; i++ {
		days = append(days, date.Format(DATE_FORMAT))
		date = AddDay(1, date)
	}
	return days
}

// 获取num个天的日期数
func getNumDays(num int, date time.Time) []string {
	var days []string
	for i := 0; i < num; i++ {
		days = append(days, date.Format(DATE_FORMAT))
		date = AddDay(1, date)
	}
	return days
}

// 获取这几个月
func GetMonth(num int, date time.Time) []string {
	var months []string
	for i := 0; i < num; i++ {
		months = append(months, date.Format(DATE_FORMAT_YYYYMM))
		date = AddMonth(1, date)
	}
	return months
}

// 获取下几个月 是几月
func GetCurrentDateNextNumMonthNum(offset int, date time.Time) (year, month int) {
	totalMonth := int(date.Month()) + offset
	month = totalMonth % 12
	currentYear := date.Year()
	year = (totalMonth/12 - 1) + currentYear
	if month == 0 {
		return year, 12
	}
	return year, month
}

// 两个日期间相差多少天,两个不同日期的，相差一秒都算一天.
// 返回昨天和今天 所以是 1 天
func DayDiffer(end, start time.Time) int {
	days := GetZeroTime(end).Sub(GetZeroTime(start)).Hours() / 24
	//return int(math.Ceil(days)) + 1
	return int(math.Ceil(days))
}

// 相差小时数
func HourDiffer(end, start time.Time) int {
	days := end.Sub(start).Hours()
	return int(math.Ceil(days))
}

// 是否是今年
func IsCurrentYear(date time.Time) bool {
	return (date.Year() - time.Now().Year()) == 0
}

// 是否是当前月份
func IsCurrentMonth(date time.Time) bool {
	return IsCurrentYear(date) && (date.Month()-time.Now().Month()) == 0
}

// 是否是今天
func IsToday(date time.Time) bool {
	return DayDiffer(date, Today()) == 0
}

// 是否是明天
func IsTomorrow(date time.Time) bool {
	return DayDiffer(date, Today()) == 1
}

// 是否是后天
func IsAfterTomorrow(date time.Time) bool {
	return DayDiffer(date, Today()) == 2
}

// 是否是大后天
func IsBigAfterTomorrow(date time.Time) bool {
	return DayDiffer(date, Today()) == 3
}

// 是否是昨天
func IsYesterday(date time.Time) bool {
	return DayDiffer(date, Today()) == -1
}

// 是否是上午 0:00 - 12:59
func IsAm(date time.Time) bool {
	hour := date.Hour()
	return hour >= 0 && hour <= 12
}

// 是否是下午 13:00 - 24:59
func IsPm(date time.Time) bool {
	hour := date.Hour()
	return hour > 12 && hour <= 24
}

//  0:00 - 12:59
func AmBetweenTime(t time.Time) (start time.Time, end time.Time) {
	start = GetZeroTime(t)
	end = GetSpecifyHourAndMinTime(t, 12, 59)
	return start, end
}

// 13:00 - 24:59
func PmBetweenTime(t time.Time) (start time.Time, end time.Time) {
	start = GetSpecifyHourAndMinTime(t, 13, 00)
	end = GetLastTime(t)
	return start, end
}

// 如果是上午 返回 00:00 - 12:59 下午 13:00 - 23:59
func StartAndEndBetweenTime(t time.Time) (start time.Time, end time.Time) {
	start = INIT_TIME
	end = INIT_TIME

	if IsAm(t) {
		start, end = AmBetweenTime(t)
	} else {
		start, end = PmBetweenTime(t)
	}
	return start, end
}

// 是否是当前小时内
func IsCurrentHour(date time.Time) bool {
	return IsToday(date) && (date.Hour()-time.Now().Hour()) == 0
}

// 两个日期间相差多少分钟,两个不同日期的
func MinuteDiffer(end, start time.Time) int {
	return int(end.Sub(start).Minutes())
}

// 离今天结束时间还有多长分钟
func TodayHasMinute() int {
	now := time.Now()
	minute := MinuteDiffer(GetLastTime(now), now)
	return minute
}

// 年维度判断两个时间是否相等.
func YearEqual(dateTime, compare time.Time) bool {
	return (dateTime.Year() - compare.Year()) == 0
}

// 月维度判断两个时间是否相等.
func MothEqual(dateTime, compare time.Time) bool {
	return YearEqual(dateTime, compare) && (dateTime.Month()-compare.Month()) == 0
}

// 天维度判断两个时间是否相等.
func DayEqual(dateTime, compare time.Time) bool {
	return MothEqual(dateTime, compare) && (dateTime.Day()-compare.Day()) == 0
}

// 小时维度判断两个时间是否相等.
func HourEqual(dateTime, compare time.Time) bool {
	return DayEqual(dateTime, compare) && (dateTime.Hour()-compare.Hour()) == 0
}

func IsInitTime(time time.Time) bool {
	return INIT_TIME.Equal(time)
}

// 在当前时间之后，就是大于当前时间
func IsNowTimeAfter(t time.Time) bool {
	return t.After(time.Now())
}

// 在当前时间之前，就是小于当前时间
func IsNowTimeBefore(t time.Time) bool {
	return t.Before(time.Now())
}

// 在当前时间之前，就是小于当前时间
func IsNowTimeEqual(t time.Time) bool {
	return t.Equal(time.Now())
}

// 在当前时间之后，就是大于当前时间（去掉秒）
func IsNowMinTimeAfter(t time.Time) bool {
	now := GetMinStartTime(time.Now())
	t = GetMinStartTime(t)
	return t.After(now)
}

// 在当前时间之前，就是小于当前时间（去掉秒）
func IsNowMinTimeBefore(t time.Time) bool {
	now := GetMinStartTime(time.Now())
	t = GetMinStartTime(t)
	return t.Before(now)
}

// 在当前时间之前，就是小于当前时间（去掉秒）
func IsNowMinTimeEqual(t time.Time) bool {
	now := GetMinStartTime(time.Now())
	t = GetMinStartTime(t)
	return t.Equal(now)
}

var weeks = []string{"一", "二", "三", "四", "五", "六", "日"}

func DateCn(t time.Time) string {
	/*else if IsBigAfterTomorrow(t) {
		dateCn = "大后天"
	} */
	dateCn := ""
	if IsYesterday(t) {
		dateCn = "昨天"
	} else if IsToday(t) {
		dateCn = "今天"
	} else if IsTomorrow(t) {
		dateCn = "明天"
	} else if DayDiffer(time.Now(), t) < 7 {
		week := GetWeek(t)
		dateCn = WeekCn(week)
	} else if DayDiffer(time.Now(), t) < 365 {
		dateCn = t.Format(DATE_FORMAT_MONTH_DAY_CN)
	} else {
		dateCn = t.Format(DATE_FORMAT_DATE_CN)
	}
	return dateCn
}

func WeekCn(week int) string {
	return fmt.Sprintf("周%s", weeks[week-1])
}

// 格式话成本地时间
func ParseLocalTime(layout, value string) time.Time {
	dateTime, err := time.ParseInLocation(layout, value, LOCAL_TIME)
	if err != nil {
		log.Println("parse time err", err)
		return INIT_TIME
	}
	return dateTime
}

func ParseLocalTimeError(layout, value string) (time.Time, error) {
	dateTime, err := time.ParseInLocation(layout, value, LOCAL_TIME)
	if err != nil {
		log.Println("parse time err", err)
		return INIT_TIME, err
	}
	return dateTime, nil
}

func FormatTime(time time.Time, layout string) string {
	timeCn := ""
	if !IsInitTime(time) {
		timeCn = time.Format(layout)
	}
	return timeCn
}

// 时间戳 to 时间
func UnixToTime(sec string) (time.Time, error) {
	data, err := strconv.ParseInt(sec, 10, 64)
	datatime := time.Unix(data/1000, 0)
	return datatime, err
}
