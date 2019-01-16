package gotime

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//GoTime 基本结构
type GoTime struct {
	time.Time
}

//ParseMap 日期格式
var ParseMap = map[string]string{
	"yyyy": "2006",
	"MM":   "01",
	"dd":   "02",
	"HH":   "15",
	"mm":   "04",
	"ss":   "05",
}

//New 初始化
func New(t time.Time) *GoTime {
	return &GoTime{t}
}

//NewUnix 将时间戳秒转为日期
func NewUnix(timestamp int64) *GoTime {
	t := time.Unix(timestamp, 0)
	return New(t)
}

//NewUnixNano 将纳秒时间戳转为日期
func NewUnixNano(timestamp int64) *GoTime {
	t := time.Unix(timestamp / 1e9, 0)
	return New(t)
}

//NewParse 初始化解析字符串日期
func NewParse(layout string, value string, timezone string) *GoTime {
	if timezone == "" {
		timezone = "Local"
	}
	for k, v := range ParseMap {
		layout = strings.Replace(layout, k, v, -1)
	}
	loc, _ := time.LoadLocation(timezone)
	t, _ := time.ParseInLocation(layout, value, loc)
	return &GoTime{t}
}

//Now 返回当前日期
func Now() (now time.Time) {
	now = time.Now()
	return
}

//Yeaterday 返回昨天日期
func Yeaterday() (yesterday *GoTime) {
	yesterday = New(time.Now()).AddDays(-1)
	return
}

//Tomorrow 返回明天日期
func Tomorrow() (tomorror *GoTime) {
	tomorror = New(time.Now()).AddDays(1)
	return
}

//Today 返回今日时间
func Today() *GoTime {
	dateStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", dateStr)
	return &GoTime{t}
}

//Parse 日期格式化
func (goTime *GoTime) Parse(format string) (parse string, err error) {
	// 解析格式
	err = nil
	for k, v := range ParseMap {
		format = strings.Replace(format, k, v, -1)
	}
	parse = goTime.Format(format)
	return
}

//AddHours 加减小时
func (goTime *GoTime) AddHours(number int) (res *GoTime) {
	m, _ := time.ParseDuration(strconv.Itoa(number) + "h")
	date := goTime.Add(m)
	res = &GoTime{date}
	return
}

//AddDays 加减天数
func (goTime *GoTime) AddDays(number int) (res *GoTime) {
	date := goTime.AddDate(0, 0, number)
	res = &GoTime{date}
	return
}

//AddMonths 加减月数
func (goTime *GoTime) AddMonths(number int) (res *GoTime) {
	date := goTime.AddDate(0, number, 0)
	res = &GoTime{date}
	return
}

//AddYears 加减年数
func (goTime *GoTime) AddYears(number int) (res *GoTime) {
	date := goTime.AddDate(number, 0, 0)
	res = &GoTime{date}
	return res
}

//toDateTimeString 转为字符串
func (goTime *GoTime) toDateTimeString() (res string) {
	res = goTime.String()
	return
}

//TimeZone 转换时区
func (goTime *GoTime) TimeZone(zone string) (res *GoTime) {
	loc,_ := time.LoadLocation(zone)
	res = &GoTime{goTime.In(loc)}
	return
}

//Timestamp 日期转为时间戳
func (goTime *GoTime) Timestamp() (timestamp int64) {
	timestamp = goTime.Unix()
	return timestamp
}

//Get 获取指定时间单位
func (goTime *GoTime) Get(dimension string) (res int) {
	switch dimension {
	case "year":
		return goTime.GetYear()
	case "month":
		return goTime.GetMonth()
	case "day":
		return goTime.GetDay()
	case "hour":
		return goTime.GetHour()
	case "minute":
		return goTime.GetMinute()
	case "second":
		return goTime.GetSecond()
	case "nanosecond":
		return goTime.GetNanosecond()
	default:
		return 0
	}
}

//GetYear 获取年份
func (goTime *GoTime) GetYear() (res int) {
	res = goTime.Year()
	return
}

//Months 月份字典
var Months = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

//GetMonth 获取月份
func (goTime *GoTime) GetMonth() (res int) {
	res = Months[goTime.Month().String()]
	return
}

//GetHour 获取小时
func (goTime *GoTime) GetHour() (res int) {
	res = goTime.Hour()
	return
}

//GetMinute 获取分钟
func (goTime *GoTime) GetMinute() (res int) {
	res = goTime.Minute()
	return
}

//GetSecond 获取秒数
func (goTime *GoTime) GetSecond() (res int) {
	res = goTime.Second()
	return
}

//GetNanosecond 获取纳秒
func (goTime *GoTime) GetNanosecond() (res int) {
	res = goTime.Nanosecond()
	return
}

//GetDay 获取天数
func (goTime *GoTime) GetDay() (res int) {
	res = goTime.Day()
	return

}

//DateSetDimension 获取日期集合
var DateSetDimension = map[string]int8{
	"month": 1,
	"day":   1,
	"hour":  1,
}

//GetDateSetFromTimestamp 计算日期集合
// @params startTime 开始时间戳
// @params startTime 结束时间戳
// @params dimension 日期间隔 如:day/hour
// @params format 日期格式 如:yyyy-MM-dd HH ...
// @params timeZone  时区
// @res dateSet 日期切片集合
// @res err 错误消息
func GetDateSetFromTimestamp(startTime int64, endTime int64, dimension string, format string, timezone string) (dateSet []string, err error) {
	if ok := DateSetDimension[dimension]; ok == 0 {
		err = errors.New("日期间隔只允许month, day, hour类型")
		return
	}
	if timezone == "" {
		timezone = "Local"
	}
	if startTime > endTime {
		err = errors.New("开始时间戳大于结束时间")
		return
	}
	startDate := New(time.Unix(startTime, 0)).TimeZone(timezone)

	// 获取格式化后的结束时间戳
	endDate, err := New(time.Unix(endTime, 0)).TimeZone(timezone).Parse(format)
	endTimestamp := NewParse(format, endDate, timezone).Unix()

	// 计算日期集合
	for {
		dateStr, _ := startDate.Parse(format)

		// 获取格式化后的累加时间戳
		startDateParse, _ := startDate.TimeZone(timezone).Parse(format)
		startTimestamp := NewParse(format, startDateParse, timezone).Unix()

		// 比较日期大小
		if startTimestamp > endTimestamp {
			break
		} else {
			dateSet = append(dateSet, dateStr)
		}
		if dimension == "month" {
			startDate = startDate.AddMonths(1)
		} else if dimension == "day" {
			startDate = startDate.AddDays(1)
		} else if dimension == "hour" {
			startDate = startDate.AddHours(1)
		} else {
			break
		}
	}
	return dateSet, nil
}
