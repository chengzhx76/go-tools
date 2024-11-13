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

	offsetAfterMonthFirstDate := StartOfMonth(dateTime).Local().AddDate(0, offset, 0)
	offsetAfterMonthLastDate := EndOfMonth(offsetAfterMonthFirstDate)
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
	offsetAfterMonthFirstDate := StartOfMonth(dateTime).Local().AddDate(offset, 0, 0)
	offsetAfterMonthLastDate := EndOfMonth(offsetAfterMonthFirstDate)

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
// Deprecated 建议使用 StartOfDay
func GetZeroTime(dateTime time.Time) time.Time {
	return StartOfDay(dateTime)
}

// 获取某一天的0点时间
func StartOfDay(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, dateTime.Location())
}

// 获取某一天指定的时间
func GetSpecifyHourAndMinTime(dateTime time.Time, hour, min int) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), hour, min, 0, 0, dateTime.Location())
}

// 获取某一分钟的0秒时间
// Deprecated 建议使用 StartOfMin
func GetMinStartTime(dateTime time.Time) time.Time {
	return StartOfMin(dateTime)
}
func StartOfMin(dateTime time.Time) time.Time {
	return GetSpecifyHourAndMinTime(dateTime, dateTime.Hour(), dateTime.Minute())
}

// 获取某一天的最后时间
// Deprecated 建议使用 EndOfDay
func GetLastTime(dateTime time.Time) time.Time {
	return EndOfDay(dateTime)
}

// 获取某一天的最后时间
func EndOfDay(dateTime time.Time) time.Time {
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 23, 59, 59, 0, dateTime.Location())
}

// 获取某一天的 开始时间和结束时间
// Deprecated 建议使用 StartAndEndOfDay
func GetDayBetweenTime(dateTime time.Time) (time.Time, time.Time) {
	return StartAndEndOfDay(dateTime)
}
func StartAndEndOfDay(dateTime time.Time) (time.Time, time.Time) {
	return StartOfDay(dateTime), EndOfDay(dateTime)
}

// 获取当前日期的年份
// Deprecated 建议使用 Year
func GetYear(date time.Time) int {
	return Year(date)
}
func Year(date time.Time) int {
	return date.Year()
}

// 返回当前日期是星期几
// Deprecated 建议使用 Week
func GetWeek(dateTime time.Time) int {
	return Week(dateTime)
}
func Week(dateTime time.Time) int {
	week := int(dateTime.Weekday())
	if week == 0 {
		return 7 // 周日改为 7
	}
	return week
}

// 返回当前日期是之后一周的时间
// Deprecated 建议使用 WeekDays
func GetWeekDays(dateTime time.Time, format string) []string {
	return WeekDays(dateTime, format)
}
func WeekDays(dateTime time.Time, format string) []string {
	days := []string{
		dateTime.Format(format),
		AddDay(1, dateTime).Format(format),
		AddDay(2, dateTime).Format(format),
		AddDay(3, dateTime).Format(format),
		AddDay(4, dateTime).Format(format),
		AddDay(5, dateTime).Format(format),
		AddDay(6, dateTime).Format(format),
	}

	return days
}

// 获取当前日期是几号
// Deprecated 建议使用 Day
func GetDay(date time.Time) int {
	return Day(date)
}
func Day(date time.Time) int {
	return date.Day()
}

// 获取从当前 date 算有之后的 cycle 个月有几天；指定日期 date 后经过几个月有几天
func GetMonthDayNums(cycle int, date time.Time) int {
	days := 0
	for i := 0; i < cycle; i++ {
		date = AddMonth(1, date)
		days += EndOfMonth(date).Day()
	}
	return days
}

// dateTime 距离 weekNum（周几）还有几天
func FromWeekDays(weekNum int, dateTime time.Time) int {
	todayWeek := Week(dateTime)
	offset := 0
	if weekNum >= todayWeek {
		offset = weekNum - todayWeek
	} else {
		offset = 7 - todayWeek + weekNum
	}
	return offset
}

// 获取某年的第一天0点时间
func YearStartDay(year string) time.Time {
	if len(year) != 4 {
		return INIT_TIME
	}
	return StartOfYear(ParseLocalTime(DATE_FORMAT_YYYYMMDD, fmt.Sprintf("%s0101", year)))
}

// 获取某年的最后一天23:59点时间
func YearEndDay(year string) time.Time {
	if len(year) != 4 {
		return INIT_TIME
	}
	return EndOfYear(ParseLocalTime(DATE_FORMAT_YYYYMMDD, fmt.Sprintf("%s1231", year)))
}

// 获取某月的第一天0点时间
func MonthStartDay(month string) time.Time {
	if len(month) != 6 {
		return INIT_TIME
	}
	return StartOfMonth(ParseLocalTime(DATE_FORMAT_YYYYMMDD, fmt.Sprintf("%s01", month)))
}

// 获取某月的最后一天23:59点时间
func MonthEndDay(month string) time.Time {
	if len(month) != 6 {
		return INIT_TIME
	}
	return EndOfMonth(ParseLocalTime(DATE_FORMAT_YYYYMMDD, fmt.Sprintf("%s31", month)))
}

// 获取某天的第一天0点时间
func DayStartDay(day string) time.Time {
	if len(day) != 8 {
		return INIT_TIME
	}
	return StartOfMonth(ParseLocalTime(DATE_FORMAT_YYYYMMDD, day))
}

// 获取某天的最后一天23:59点时间
func DayEndDay(day string) time.Time {
	if len(day) != 8 {
		return INIT_TIME
	}
	return EndOfMonth(ParseLocalTime(DATE_FORMAT_YYYYMMDD, day))
}

// 获取某年的第一天0点时间
func StartOfYear(year time.Time) time.Time {
	return time.Date(year.Year(), time.January, 1, 0, 0, 0, 0, year.Location())
}

// 获取某年的最后一天23:59:59点时间
func EndOfYear(year time.Time) time.Time {
	return time.Date(year.Year(), time.December, 31, 23, 59, 59, 0, year.Location())
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
// Deprecated 建议使用 StartOfMonth
func GetFirstDateOfMonth(d time.Time) time.Time {
	//d = d.AddDate(0, 0, -d.Day()+1)
	//return GetZeroTime(d)
	return StartOfMonth(d)
}
func StartOfMonth(month time.Time) time.Time {
	return time.Date(month.Year(), month.Month(), 1, 0, 0, 0, 0, month.Location())
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的23:59:59点。如传入time.Now(), 返回当前月份的最后一天23:59:59点时间。
// Deprecated 建议使用 EndOfMonth
func GetLastDateOfMonth(d time.Time) time.Time {
	//return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
	return EndOfMonth(d)
}
func EndOfMonth(month time.Time) time.Time {
	return time.Date(month.Year(), month.Month(), 31, 23, 59, 59, 0, month.Location())
}

// 获取num个月的日期数
func GetNumMonthDays(num int, date time.Time) []string {
	differDays := DayDiffer(EndOfMonth(AddMonth(num, date)), date)
	var days []string
	var i int
	for i = 0; i < differDays; i++ {
		days = append(days, date.Format(DATE_FORMAT))
		date = AddDay(1, date)
	}
	return days
}

// 获取num个天的日期数
// Deprecated 建议使用 NumDays
func GetNumDays(num int, date time.Time) []string {
	return NumDays(num, date)
}
func NumDays(num int, date time.Time) []string {
	var days []string
	for i := 0; i < num; i++ {
		days = append(days, date.Format(DATE_FORMAT))
		date = AddDay(1, date)
	}
	return days
}

// 从date开始获取这num月的字符串
// Deprecated 建议使用 NumMonths
func GetMonth(num int, date time.Time) []string {
	return NumMonths(num, date)
}
func NumMonths(num int, date time.Time) []string {
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

// 两个日期间相差多少天,两个不同日期的，不同天比较，相同天是0，负数表示开始时间大约结束时间.
// 返回昨天和今天 所以是 1 天；
func DayDiffer(end, start time.Time) int {
	days := StartOfDay(end).Sub(StartOfDay(start)).Hours() / 24
	//return int(math.Ceil(days)) + 1
	return int(math.Ceil(days))
}

// 两个不同日期的，相差一秒都算一天. 两个日期相等就是0
func DaySecondDiffer(end, start time.Time) int {
	if end.Equal(start) {
		return 0
	} else {
		seconds := EndOfDay(end).Sub(EndOfDay(start)).Seconds()
		return If[int](seconds == 0, 1, int(seconds/(24*60*60)))
	}
}

// 相差小时数
func HourDiffer(end, start time.Time) int {
	days := end.Sub(start).Hours()
	return int(math.Ceil(days))
}

// 两个日期间相差多少分钟,两个不同日期的
func MinuteDiffer(end, start time.Time) int {
	return int(end.Sub(start).Minutes())
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

// 0:00 - 12:59
func AmBetweenTime(t time.Time) (start time.Time, end time.Time) {
	start = StartOfDay(t)
	end = GetSpecifyHourAndMinTime(t, 12, 59)
	return start, end
}

// 13:00 - 24:59
func PmBetweenTime(t time.Time) (start time.Time, end time.Time) {
	start = GetSpecifyHourAndMinTime(t, 13, 00)
	end = EndOfDay(t)
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
	return
}

// 偏移offset分钟
func OffsetStartAndEndMinutes(t time.Time, offset int64) (start time.Time, end time.Time) {
	start = AddMinute(-offset, t)
	end = AddMinute(offset, t)
	return
}

// 是否是当前小时内
func IsCurrentHour(date time.Time) bool {
	return IsToday(date) && (date.Hour()-time.Now().Hour()) == 0
}

// 离今天结束时间还有多长分钟
func TodayHasMinute() int {
	now := time.Now()
	minute := MinuteDiffer(EndOfDay(now), now)
	return minute
}

// 现在是今天的第 %d 分钟
func MinutesToday() int {
	now := time.Now()
	return now.Hour()*60 + now.Minute()
}

// 分钟转小时
func MinutesToHours(minutes int) int {
	return minutes / 60
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

// 分钟维度判断两个时间是否相等.
func MinEqual(dateTime, compare time.Time) bool {
	return HourEqual(dateTime, compare) && (dateTime.Minute()-compare.Minute()) == 0
}

// 判断 checkTime 是否在 startTime 和 endTime 时间之间
func TimeBetween(checkTime, startTime, endTime time.Time) bool {
	return (MinEqual(checkTime, startTime) || MinEqual(checkTime, endTime)) || checkTime.After(startTime) && checkTime.Before(endTime)
}

// 是否是初始时间
func IsInitTime(time time.Time) bool {
	return INIT_TIME.Equal(time)
}

// 是否是最终时间
func IsEndLessTime(time time.Time) bool {
	return END_LESS_TIME.Equal(time)
}

// 在当前时间之后（就是未来的时间），就是大于当前时间
func IsNowTimeAfter(t time.Time) bool {
	return t.After(time.Now())
}

// 在当前时间之前（就是已过去的时间），就是小于当前时间
func IsNowTimeBefore(t time.Time) bool {
	return t.Before(time.Now())
}

// 等于当前时间
func IsNowTimeEqual(t time.Time) bool {
	return t.Equal(time.Now())
}

// 结束时间在开始时间之后，就是`end`大于`start`时间（去掉秒）
func EndMinTimeAtStartMinTimeAfter(start, end time.Time) bool {
	start = StartOfMin(start)
	end = StartOfMin(end)
	return end.After(start)
}

// 在当前时间之后（就是未来的时间），就是大于当前时间（去掉秒）
func IsNowMinTimeAfter(t time.Time) bool {
	now := StartOfMin(time.Now())
	t = StartOfMin(t)
	return t.After(now)
}

// 在当前时间之前（就是已过去的时间），就是小于当前时间（去掉秒）
func IsNowMinTimeBefore(t time.Time) bool {
	now := StartOfMin(time.Now())
	t = StartOfMin(t)
	return t.Before(now)
}

// 等于当前时间（去掉秒）
func IsNowMinTimeEqual(t time.Time) bool {
	now := StartOfMin(time.Now())
	t = StartOfMin(t)
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
		week := Week(t)
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
	datetime := time.Unix(data/1000, 0)
	return datetime, err
}

type TimeFormatDTO struct {
	Num     int
	Uint    string
	Tips    string
	Expired bool
}

func TimeFormat(t time.Time) *TimeFormatDTO {
	now := time.Now()
	diffMin := MinuteDiffer(t, now)
	expired := diffMin < 0
	suffix := "后"
	if expired {
		suffix = "前"
	}
	diffMin = int(math.Abs(float64(diffMin)))
	if diffMin > 60*24*31 {
		monthDay := int(math.Floor(float64(diffMin) / float64(60*24*31)))
		return &TimeFormatDTO{
			Num:     monthDay,
			Uint:    "月",
			Tips:    fmt.Sprintf("%d月%s", monthDay, suffix),
			Expired: expired,
		}
	} else if diffMin > 60*24 {
		diffDay := int(math.Floor(float64(diffMin) / float64(60*24)))
		return &TimeFormatDTO{
			Num:     diffDay,
			Uint:    "天",
			Tips:    fmt.Sprintf("%d天%s", diffDay, suffix),
			Expired: expired,
		}
	} else if diffMin > 60 {
		diffHour := int(math.Floor(float64(diffMin) / float64(60)))
		if diffHour == 24 {
			return &TimeFormatDTO{
				Num:     1,
				Uint:    "天",
				Tips:    fmt.Sprintf("1天%s", suffix),
				Expired: expired,
			}
		} else {
			return &TimeFormatDTO{
				Num:     diffHour,
				Uint:    "时",
				Tips:    fmt.Sprintf("%d小时%s", diffHour, suffix),
				Expired: expired,
			}
		}
	} else { // 小于等于1小时
		if diffMin == 60 {
			return &TimeFormatDTO{
				Num:     1,
				Uint:    "时",
				Tips:    fmt.Sprintf("1小时%s", suffix),
				Expired: expired,
			}
		} else if diffMin > 3 {
			return &TimeFormatDTO{
				Num:     diffMin,
				Uint:    "分",
				Tips:    fmt.Sprintf("%d分钟%s", diffMin, suffix),
				Expired: expired,
			}
		} else {
			return &TimeFormatDTO{
				Num:     diffMin,
				Uint:    "分",
				Tips:    "刚刚",
				Expired: expired,
			}
		}
	}
}
