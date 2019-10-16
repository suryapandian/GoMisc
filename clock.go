package main

import "strconv"
import "fmt"

type Clock struct{ h, m int }

func main() {
	New(0, 45).Add(40)
}

func New(hour, minute int) Clock {
	i := Clock{hour, minute}
	i.correctTime()
	return i
}

func (i Clock) String() (result string) {
	i.correctTime()
	minute := i.getMinute()
	hour := i.getHour()
	result = hour + ":" + minute
	return
}

func (i Clock) Add(minutes int) Clock {
	result := Clock{i.h, i.m + minutes}
	fmt.Println("Add:", result.m, result, result.m)
	return result
}

func (i Clock) Subtract(minutes int) Clock {
	result := Clock{i.h, i.m - minutes}
	return result
}

func (i Clock) validate() bool {
	if (i.m >= 0) && (i.h > 0) {
		return true
	}
	return false
}

func (i *Clock) correctTime() {
	hours, mins := i.m/60, i.m%60
	i.h = i.h + hours
	i.m = mins
	i.absoluteHour()
	i.absoluteMin()
}

func (i *Clock) absoluteHour() {
	i.h = absolute(i.h, 24)
}

func absolute(i int, n int) int {
	var negative bool
	if i < 0 {
		negative = true
		i = i * (-1)
	}
	i = i % n
	if negative {
		i = n - i
	}
	return i
}

func (i *Clock) absoluteMin() {
	if i.m < 0 {
		if i.h > 0 {
			i.h = i.h - 1
		} else {
			i.h = 23
		}
	}
	i.m = absolute(i.m, 60)
}

func (i Clock) getMinute() (result string) {
	result = appendZero(strconv.Itoa(i.m % 60))
	return
}

func (i Clock) getHour() (result string) {
	if i.h == 24 {
		i.h = 0
	}
	result = strconv.Itoa(i.h % 24)
	result = appendZero(result)
	return
}

func appendZero(time string) (result string) {
	result = time
	if len(time) == 1 {
		result = "0" + time
	}
	return
}
