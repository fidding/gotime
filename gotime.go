package gotime

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type GoTime struct {
	time.Time
}

// 日期格式
var ParseMap = map[string]string{
	"yyyy": "2006",
	"MM":   "01",
	"dd":   "02",
	"HH":   "15",
	"mm":   "04",
	"ss":   "05",
}

// 初始化
func New(t time.Time) *GoTime {
	return &GoTime{t}
}

// 初始化解析字符串日期
func NewParse(layout string, value string) *GoTime {
	for k, v := range ParseMap {
		layout = strings.Replace(layout, k, v, -1)
	}
	t, _ := time.Parse(layout, value)
	return &GoTime{t}
}

// 返回当前日期
func Now() (now time.Time) {
	now = time.Now()
	return
}

// 返回昨天日期
func Yeaterday() (yesterday *GoTime) {
	yesterday = New(time.Now()).AddDays(-1)
	return
}

// 返回明天日期
func Tomorrow() (tomorror *GoTime) {
	tomorror = New(time.Now()).AddDays(1)
	return
}

// 返回今日时间
func Today() *GoTime {
	dateStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", dateStr)
	return &GoTime{t}
}

// 日期格式化
func (goTime *GoTime) Parse(format string) (parse string, err error) {
	// 解析格式
	err = nil
	for k, v := range ParseMap {
		format = strings.Replace(format, k, v, -1)
	}
	parse = goTime.Format(format)
	return
}

// 加减小时
func (goTime *GoTime) AddHours(number int) (res *GoTime) {
	m, _ := time.ParseDuration(strconv.Itoa(number) + "h")
	date := goTime.Add(m)
	res = &GoTime{date}
	return
}

// 加减天数
func (goTime *GoTime) AddDays(number int) (res *GoTime) {
	date := goTime.AddDate(0, 0, number)
	res = &GoTime{date}
	return
}

// 加减月数
func (goTime *GoTime) AddMonths(number int) (res *GoTime) {
	date := goTime.AddDate(0, number, 0)
	res = &GoTime{date}
	return
}

// 加减年数
func (goTime *GoTime) AddYears(number int) (res *GoTime) {
	date := goTime.AddDate(number, 0, 0)
	res = &GoTime{date}
	return res
}

// 转为字符串
func (goTime *GoTime) toDateTimeString() (res string) {
	res = goTime.String()
	return
}

// 转换时区
func (goTime *GoTime) TimeZone(zone string) (res *GoTime) {
	loc, _ := time.LoadLocation(zone)
	res = &GoTime{goTime.In(loc)}
	return
}

// 日期转为时间戳
func (goTime *GoTime) Timestamp() (timestamp int64) {
	timestamp = goTime.Unix()
	return timestamp
}

// 获取日期集合
var DateSetDimension = map[string]int8{
	"month": 1,
	"day": 1,
	"hour": 1,
}

// 计算日期集合
// @params startTime 开始时间戳
// @params startTime 结束时间戳
// @params dimension 日期间隔 如:day/hour
// @params format 日期格式 如:yyyy-MM-dd HH ...
// @params timeZone  时区
// @res dateSet 日期切片集合
// @res err 错误消息
func GetDateSetFromTimestamp(startTime int64, endTime int64, dimension string, format string, timeZone string) (dateSet []string, err error){
	if ok := DateSetDimension[dimension]; ok == 0 {
		err =  errors.New("日期间隔只允许month, day, hour类型")
		return
	}
	if timeZone == "" {
		timeZone = "Local"
	}
	if startTime > endTime {
		err = errors.New("开始时间戳大于结束时间")
		return
	}
	startDate := New(time.Unix(startTime, 0)).TimeZone(timeZone)

	// 获取格式化后的结束时间戳
	endDate, err := New(time.Unix(endTime, 0)).TimeZone(timeZone).Parse(format)
	endTimestamp := NewParse(format, endDate).TimeZone(timeZone).Unix()

	// 计算日期集合
	for {
		dateStr, _ := startDate.Parse(format)

		// 获取格式化后的累加时间戳
		startDateParse, _ := startDate.TimeZone(timeZone).Parse(format)
		startTimestamp := NewParse(format, startDateParse).TimeZone(timeZone).Unix()

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
		} else  if dimension == "hour" {
			startDate = startDate.AddHours(1)
		} else {
			break
		}
	}

	return dateSet, nil
}
