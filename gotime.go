package gotime

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type GoTime struct{
	time.Time
}

// 日期格式
var ParseMap = map[string]string{
	"yyyy": "2006",
	"MM": "01",
	"dd": "02",
	"HH": "15",
	"mm": "04",
	"ss": "05",
}

// New initialize Now with time
func New(t time.Time) *GoTime {
	return &GoTime{t}
}

// 返回当前时间
func Now() time.Time {
	now := time.Now()
	return now
}

// 返回昨天时间
func Yeaterday() *GoTime {
	now := time.Now()
	yesterday := New(now).AddDays(-1)
	return yesterday
}

// 返回今日时间
func Today() *GoTime {
	dateStr := time.Now().Format("2006-01-02")
	t, _ :=  time.Parse("2006-01-02", dateStr)
	return &GoTime{t}
}

// 日期格式化
func (goTime *GoTime) Parser(parser string) (string, error) {
	// 解析格式
	for format, parse := range ParseMap {
		parser = strings.Replace(parser, format, parse, -1)
	}
	return  goTime.Format(parser), nil
}

// 加减小时
func (goTime *GoTime) AddHours(number int) *GoTime {
	m, _ := time.ParseDuration(strconv.Itoa(number) + "h")
	date := goTime.Add(m)
	fmt.Println(strconv.Itoa(number) + "h")
	res := &GoTime{date}
	return res
}

// 加减天数
func (goTime *GoTime) AddDays(number int) *GoTime {
	date := goTime.AddDate(0, 0, number)
	res := &GoTime{date}
	return res
}

// 加减月数
func (goTime *GoTime) AddMonths(number int) *GoTime {
	date := goTime.AddDate(0, number, 0)
	res := &GoTime{date}
	return res
}

// 加减年数
func (goTime *GoTime) AddYears(number int) *GoTime {
	date := goTime.AddDate(number, 0, 0)
	res := &GoTime{date}
	return res
}

// 转为字符串
func (goTime *GoTime) toDateTimeString() string {
	return goTime.String()
}

// 转换时区
func (goTime *GoTime) TimeZone(zone string) *GoTime {
	lo, _ := time.LoadLocation(zone)
	return &GoTime{goTime.In(lo)}
}