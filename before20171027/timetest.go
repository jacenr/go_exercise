package main

import (
	"fmt"
	"time"
)

const shortForm = "2006-01-02 15:04 MST"
const longForm = "Mon Jan 2 15:04:05 -0700 MST 2006"

func main() {
	timeNow := time.Now()

	// 获取一个月前或者一年前的时间
	aMonthAgo := time.Date(timeNow.Year(), timeNow.Month()-1, timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), timeNow.Location())
	aYear := time.Date(timeNow.Year()-1, timeNow.Month(), timeNow.Day(), timeNow.Hour(), timeNow.Minute(), timeNow.Second(), timeNow.Nanosecond(), timeNow.Location())
	fmt.Println(aMonthAgo)
	fmt.Println(aYear)

	// 额外测试: 今天具体上个月同一时间的间隔小时数
	fmt.Println(time.Since(aMonthAgo).Hours())

	// 时间前后测试, 返回bool, Befor, After, Equal
	fmt.Println(timeNow.Before(aMonthAgo))

	// 获取指定日期的时间
	loc, err := time.LoadLocation("Local")
	if err != nil {
		fmt.Printf("load location err: %v", err)
	}
	t1 := time.Date(2017, 07, 10, 10, 30, 30, 1000, loc)
	fmt.Println(t1)

	// 从字符串解析时间
	t2, err := time.Parse(shortForm, "2017-07-10 10:30 CST") // 格式1
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)
	t3, err := time.Parse(longForm, "Mon Jul 10 10:30:05 +0800 CST 2017") // 格式2
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3)

	// 利用时间差获取时间
	dur, _ := time.ParseDuration("-10h")
	t4 := timeNow.Add(dur)
	fmt.Println(t4)
}
