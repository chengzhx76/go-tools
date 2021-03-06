package consts

import (
	"time"
)

var (
	INIT_TIME        = time.Unix(0, 0)
	END_LESS_TIME, _ = time.ParseInLocation(DATE_FORMAT_YYYYMMDDHHMMSS, "99991231235959", LOCAL_TIME)
	LOCAL_TIME, _    = time.LoadLocation("Local")
)

const (
	SYMBOL_DOT           = "."
	SYMBOL_ASTERISK      = "*"
	SYMBOL_COMMA         = ","
	SYMBOL_HASHTAG       = "#"
	SYMBOL_VERTICAL      = "|"
	SYMBOL_EMPTY         = ""
	SYMBOL_SPACE         = " "
	SYMBOL_COLON         = ":"
	SYMBOL_UNDERLINE     = "-"
	SYMBOL_AT            = "@"
	ERROR_CODE_DELIMITER = "__"
)

const (
	DATE_FORMAT_MONTH               = "2006-01"
	DATE_FORMAT                     = "2006-01-02"
	DATE_FORMAT_DATE_CN             = "2006年01月02日"
	DATE_FORMAT_MONTH_DAY_CN        = "01月02日"
	DATE_FORMAT_DATE_TIME_CN        = "01月02日 15:04"
	DATE_FORMAT_TIME_CN             = "15:04:05"
	DATE_FORMAT_HOUR_MINUTE_CN      = "15:04"
	DATE_FORMAT_MINUTE_CN           = "2006年01月02日 15:04"
	DATE_TIME_FORMAT                = "2006-01-02 15:04:05"
	DATE_TIME_MINUTE_FORMAT         = "2006-01-02 15:04"
	DATE_TIME_MINUTE_FORMAT_NOCOLON = "2006-01-02 1504"
	DATE_FORMAT_YYYYMM              = "200601"         // 精确到月
	DATE_FORMAT_YYYYMMDD            = "20060102"       // 精确到天
	DATE_FORMAT_YYYYMMDDHH          = "2006010215"     // 精确到小时
	DATE_FORMAT_YYYYMMDDHHMM        = "200601021504"   // 精确到分
	DATE_FORMAT_YYYYMMDDHHMMSS      = "20060102150405" // 精确到秒
	DATE_FORMAT_HHMM                = "1504"           // 时分
	TIME_FORMAT                     = "15:04:05"
)

const (
	TRUE  bool = true
	FALSE bool = false
)

const (
	YES uint8 = 1
	NO  uint8 = 0
)

const (
	DELETED uint8 = 0 // 删除
	NORMAL  uint8 = 1 // 正常
)

const (
	HIDE uint8 = 0 // 隐藏
	SHOW uint8 = 1 // 显示
)

const (
	ADD uint8 = 1 // 增加
	SUB uint8 = 2 // 减少
)

const (
	WEEKDAY uint8 = 0 // 工作日
	WEEKEND uint8 = 1 // 假日（周末）
	HOLIDAY uint8 = 2 // 节假日
)

// 货币类型
const (
	CURRENCY_TYPE_CNY string = "CNY" // 人民币
)

const (
	HOLIDAYS_YUANDAN  uint8 = 1
	HOLIDAYS_CHUNJIE  uint8 = 2
	HOLIDAYS_QINGMING uint8 = 3
	HOLIDAYS_LAODONG  uint8 = 4
	HOLIDAYS_DUANWU   uint8 = 5
	HOLIDAYS_ZHONGQIU uint8 = 6
	HOLIDAYS_GUOQING  uint8 = 7
)
