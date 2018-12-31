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
#### 日期格式化
```go
gotime, err := gotime.New(time.Now()).Parse("yyyy-MM-dd HH:mm:ss")

```

#### 时区转换
```go
date := gotime.New(time.Now()).TimeZone("Local")
```

#### 日期加减
```go
date := gotime.New(time.Now()).AddYears(1)
date := gotime.New(time.Now()).AddDays(-1)
date := gotime.New(time.Now()).AddHours(1)
```

#### 日期切片计算
```go
// 返回["2018-12-30 00", "2018-12-30 01", "2018-12-30 02", "2018-12-30 03", "2018-12-30 04"]
gotime.GetDateSetFromTimestamp(1546099200, 1546114200, "hour", "yyyy-MM-dd HH", "")
```

# Author

* [http://www.fidding.me](http://www.fidding.me)
* [http://github.com/fidding](http://github.com/fidding)

# License
Released under the [MIT License](http://www.opensource.org/licenses/MIT).