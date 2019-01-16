package gotime

import (
	"fmt"
	"testing"
)

var ParseMapTest = map[string]string{
	"yyyy/MM/dd HH:mm:ss": "2006/01/02 15:04:05",

	"yyyy-MM-dd HH:mm:ss": "2006-01-02 15:04:05",
	"yyyy-MM-dd HH:mm":    "2006-01-02 15:04",
	"yyyy-MM-dd HH":       "2006-01-02 15",
	"yyyy-MM-dd":          "2006-01-02",
	"yyyy-MM":             "2006-01",
	"yyyy":                "2006",

	"MM-dd HH:mm:ss": "01-02 15:04:05",
	"MM-dd HH:mm":    "01-02 15:04",
	"MM-dd HH":       "01-02 15",
	"MM-dd":          "01-02",
	"MM":             "01",

	"dd HH:mm:ss": "02 15:04:05",
	"dd HH:mm":    "02 15:04",
	"dd HH":       "02 15",
	"dd":          "02",

	"HH:mm:ss": "15:04:05",
	"HH:mm":    "15:04",
	"HH":       "15",

	"mm:ss": "04:05",
	"mm":    "04",
	"ss":    "05",
}

func TestNewParse(t *testing.T) {
	var got, want int64
	res := NewParse("yyyy-MM-dd HH", "2018-12-20 15", "Asia/Shanghai")
	got = res.Unix()
	want = 1545289200
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}

// 测试日期格式化
func TestGoTime_Parse(t *testing.T) {
	gotime := New(Now())
	for format := range ParseMapTest {
		fmt.Print(format + ": ")
		if res, err := gotime.Parse(format); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}
	}
}

func TestGoTime_AddHours(t *testing.T) {
	res := New(Now()).AddHours(-1)
	fmt.Println(res)
}

func TestGoTime_AddDays(t *testing.T) {
	res := New(Now()).AddDays(-1)
	fmt.Println(res)
}

func TestGoTime_AddMonths(t *testing.T) {
	res := New(Now()).AddMonths(-1)
	fmt.Println(res)
}

func TestGoTime_AddYears(t *testing.T) {
	res := New(Now()).AddYears(-1)
	fmt.Println(res)
}

func TestGoTime_Timestamp(t *testing.T) {
	res := New(Now()).Timestamp()
	fmt.Println(res)
}

var wantHour = []string{
	"2018-12-30 00",
	"2018-12-30 01",
	"2018-12-30 02",
	"2018-12-30 03",
	"2018-12-30 04",
}
var wantDay = []string{
	"2018-12-30",
	"2018-12-31",
}
var wantMonth = []string{
	"2018-12",
	"2019-01",
}
func TestGoTime_GetDateSetFromTimestamp(t *testing.T) {
	// 检测按小时切片
	got, err := GetDateSetFromTimestamp(1546099200, 1546114200, "hour", "yyyy-MM-dd HH", "")
	if err != nil {
		fmt.Println(err)
	}
	if len(got) != len(wantHour) {
		t.Errorf("got %v want %v given", got, wantHour)
	} else {
		for i := range got {
			if got[i] != wantHour[i] {
				t.Errorf("got %v want %v given", got, wantHour)
				break
			}
		}
	}

	// 检测按天切片
	got2, err := GetDateSetFromTimestamp(1546099200, 1546242692, "day", "yyyy-MM-dd", "")
	if err != nil {
		t.Errorf("got %v want %v given", got2, wantDay)
	}
	if len(got2) != len(wantDay) {
		t.Errorf("got %v want %v given", got2, wantDay)
	} else {
		for i := range got2 {
			if got2[i] != wantDay[i] {
				t.Errorf("got %v want %v given", got2, wantDay)
				break
			}
		}
	}

	// 检测按月切片
	got3, err := GetDateSetFromTimestamp(1546099200, 1546272000, "month", "yyyy-MM", "")
	if err != nil {
		t.Errorf("got %v want %v given", got3, wantMonth)
	}
	if len(got3) != len(wantMonth) {
		t.Errorf("got %v want %v given", got3, wantMonth)
	} else {
		for i := range got3 {
			if got3[i] != wantMonth[i] {
				t.Errorf("got %v want %v given", got3, wantMonth)
				break
			}
		}
	}
}

// 测试获取年月日时分秒获取
func TestGoTime_Get(t *testing.T) {
	testTime := NewParse("yyyy-MM-dd HH:mm:ss", "2018-02-01 10:23:59", "")
	// 年
	got := testTime.Get("year")
	want := 2018
	if got != want {
		t.Errorf("test year got %v want %v given", got, want)
	}
	// 月
	got1 := testTime.Get("month")
	want1 := 2
	if got1 != want1 {
		t.Errorf("test month got %v want %v given", got1, want1)
	}
	// 日
	got2 := testTime.Get("day")
	want2 := 1
	if got2 != want2 {
		t.Errorf("test day got %v want %v given", got2, want2)
	}
	// 时
	got3 := testTime.Get("hour")
	want3 := 10
	fmt.Println(got3, want3)
	if got3 != want3 {
		t.Errorf("test hour got %v want %v given", got3, want3)
	}
	// 分
	got4 := testTime.Get("minute")
	want4 := 23
	if got4 != want4 {
		t.Errorf("test minute got %v want %v given", got4, want4)
	}
	// 秒
	got5 := testTime.Get("second")
	want5 := 59
	if got5 != want5 {
		t.Errorf("test second got %v want %v given", got5, want5)
	}
}

//TestGoTime_CompariTo 测试日期比较
func TestGoTime_CompariTo(t *testing.T) {
	date := NewParse("yyyy-MM-dd HH:mm:ss", "1991-09-17 00:00:00", "")
	time1 := NewParse("yyyy-MM-dd HH:mm:ss", "1991-09-17 00:00:01", "")
	got1 := date.CompariTo(time1)
	want1 := -1
	if got1 != want1 {
		t.Errorf("test compari 1 got %v want %v given", want1, got1)
	}
	time2 := NewParse("yyyy-MM-dd", "1991-09-17", "")
	got2 := date.CompariTo(time2)
	want2 := 0
	if got2 != want2 {
		t.Errorf("test compari 2 got %v want %v given", want2, got2)
	}
	time3 := NewParse("yyyy/MM/dd HH:mm:ss", "1991/09/16 00:00:00", "")
	got3 := date.CompariTo(time3)
	want3 := 1
	if got3 != want3 {
		t.Errorf("test compari 3 got %v want %v given", want3, got3)
	}
}


//TestNewUnix 测试时间戳转日期
func TestNewUnix(t *testing.T) {
	var timestamp int64
	timestamp = 685036800
	got, _ := NewUnix(timestamp).Parse("yyyy-MM-dd HH:mm:ss")
	fmt.Println(got)
	want := "1991-09-17 00:00:00"
	if got != want {
		t.Errorf("test newUnix  got %v want %v given", want, got)
	}
}

//TestNewUnixNano 测试时间戳转日期
func TestNewUnixNano(t *testing.T) {
	var timestamp int64
	timestamp = 685036800 * 1e9
	got, _ := NewUnixNano(timestamp).Parse("yyyy-MM-dd HH:mm:ss")
	want := "1991-09-17 00:00:00"
	if got != want {
		t.Errorf("test newUnixNano  got %v want %v given", want, got)
	}
}
