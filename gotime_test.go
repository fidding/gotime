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


// 测试日期格式化
func TestGoTime_Parse(t *testing.T) {
	gotime := New(Now())
	for format, _ := range ParseMapTest {
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
func TestGoTime_GetDateSetFromTimestamp(t *testing.T) {
	got, err := GetDateSetFromTimestamp(1546099200, 1546114200, "hour", "yyyy-MM-dd HH","")
	if err != nil {
		fmt.Println(err)
	}
	if len(got) != len(wantHour) {
		t.Errorf("got %v want %v given", got, wantHour)
	}
	for i := range got {
		if got[i] != wantHour[i] {
			t.Errorf("got %v want %v given", got, wantHour)
			break
		}
	}

	got2, err := GetDateSetFromTimestamp(1546099200, 1546242692, "day", "yyyy-MM-dd","")
	if err != nil {
		fmt.Println(err)
	}
	if len(got2) != len(wantDay) {
		t.Errorf("got %v want %v given", got, wantHour)
	}
	for i := range got2 {
		if got2[i] != wantDay[i] {
			t.Errorf("got %v want %v given", got, wantHour)
			break
		}
	}
}