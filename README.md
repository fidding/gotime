# gotime
Is a better time  toolkit for golang

# Install

```bash
go get -u github.com/fidding/gotime
```

# Usage
```go
import "github.com/fidding/gotime"
```

#### 常用日期
```go
// 当前时间
gotime.Now()
// 今天
gotime.Today()
// 昨天
gotime.Yesterday()
// 明天
gotime.Tomorrow()
```

#### 实例化
```go
// time实例化为gotime
date := time.Now()
gotime.New(date)
// 字符串时间实例化为gotime
date := "1991-09-17"
gotime.NewParse("yyyy-MM-dd", date, "Local")
```

#### 日期获取
```go
// 获取年/月/日/时/分/秒/纳秒
// year/month/day/hour/minute/second/nanosecond
goTime.Get("year")
// or
goTime.GetYear()
goTime.GetMonth()
goTime.GetDay()
goTime.GetHour()
goTime.GetMinute()
goTime.GetSecond()
goTime.GetNanoSecond()

```

#### 日期格式化
```go
// 格式化
// yyyy 年
// MM 月
// dd 日
// HH 时
// mm 分
// ss 秒
gotime, err := gotime.New(time.Now()).Parse("yyyy-MM-dd HH:mm:ss")
```

#### 时区转换
```go
date := gotime.New(time.Now()).TimeZone("Local")
```

#### 日期加减
```go
// 年份加减
date := gotime.New(time.Now()).AddYears(1)
// 月份加减
date := gotime.New(time.Now()).AddMonths(-1)
// 天数加减
date := gotime.New(time.Now()).AddDays(-1)
// 小时加减
date := gotime.New(time.Now()).AddHours(1)
```

#### 日期比较
```go
//  res 大于1小于-1等于0
date := gotime.NewParse("yyyy-MM-dd HH:mm:ss", "1991-09-17 00:00:00", "")
date2 := gotime.NewParse("yyyy-MM-dd HH:mm:ss", "1991-09-17 00:00:01", "")
res := date.CompariTo(date2)
```

#### 日期切片计算
```go
// 返回中国时间2018-12-30 00:00:00到2018-12-30 04:10:00
// 按小时切片集合
// 结果集合["2018-12-30 00", "2018-12-30 01", "2018-12-30 02", "2018-12-30 03", "2018-12-30 04"]
gotime.GetDateSetFromTimestamp(1546099200, 1546114200, "hour", "yyyy-MM-dd HH", "Asia/Shanghai")

// 返回中国时间2018-12-30 00:00:00到2018-12-31 15:51:32
// 按天切片集合
// 结果集合["2018-12-30", "2018-12-31"]
gotime.GetDateSetFromTimestamp(1546099200, 1546242692, "day", "yyyy-MM-dd", "Asia/Shanghai")

// 返回中国时间2018-12-30 00:00:00到2019-01-01 00:00:00
// 按月切片集合
// 结果集合["2018-12", "2019-01"]
gotime.GetDateSetFromTimestamp(1546099200, 1546272000, "month", "yyyy-MM", "Asia/Shanghai")
```

# Author

* [http://www.fidding.me/about](http://www.fidding.me/about)
* [http://github.com/fidding](http://github.com/fidding)

# License
Released under the [MIT License](http://www.opensource.org/licenses/MIT).
