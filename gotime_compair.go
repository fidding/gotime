package gotime

//CompariTo 日期比较
func (goTime *GoTime) CompariTo(date *GoTime) (res int) {
	diff := goTime.UnixNano() - date.UnixNano()
	if diff == 0 {
		res = 0
	} else if diff > 0 {
		res = 1
	} else {
		res = -1
	}
	return
}
