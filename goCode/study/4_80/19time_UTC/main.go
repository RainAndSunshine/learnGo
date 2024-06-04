package main

import (
	"fmt"
	"time"
)

//时区

func main() {
	now := time.Now() //本地的时间 东八区的时间
	fmt.Println(now)
	//明天的这个时间
	//默认得出的时间是UTC的时间
	nextDay, err := time.Parse("2006-01-02 15:04:05", "2024-05-19 14:41:00")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(nextDay.Sub(now))
	//按照东八区的时区和格式去解析一个字符串类型的时间
	//根据字符串加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("load time failed, err:%v\n", err)
		return
	}
	//按照指定时区解析时间
	nextDay2, err := time.ParseInLocation("2006-01-02 15:04:05", "2024-05-19 14:41:00", loc)
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println("--------------------")
	fmt.Println("正确的时间间隔")
	fmt.Println(nextDay2.Sub(now))
}
