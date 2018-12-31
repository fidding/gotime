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

#### 日期切片计算
```go
// 返回1546099200到1546114200按小时切片集合
// 结果集合["2018-12-30 00", "2018-12-30 01", "2018-12-30 02", "2018-12-30 03", "2018-12-30 04"]
gotime.GetDateSetFromTimestamp(1546099200, 1546114200, "hour", "yyyy-MM-dd HH", "Local")

// 返回1546099200到1546242692按天切片集合
// 结果集合["2018-12-30", "2018-12-31"]
GetDateSetFromTimestamp(1546099200, 1546242692, "day", "yyyy-MM-dd", "Local")
``` 

# Author

* [http://www.fidding.me/about](http://www.fidding.me/about)
* [http://github.com/fidding](http://github.com/fidding)

# License
Released under the [MIT License](http://www.opensource.org/licenses/MIT).