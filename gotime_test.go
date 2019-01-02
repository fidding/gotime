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
