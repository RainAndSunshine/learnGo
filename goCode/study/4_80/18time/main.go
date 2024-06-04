package main

import (
	"fmt"
	"time"
)

//时间  time

func main() {
	//现在的时间
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//Unix()  时间戳
	//自1970年1月1日8:00至当前时间的总秒数
	fmt.Println(now.Unix())     //拿到当前时间的时间戳
	fmt.Println(now.UnixNano()) //纳秒时间戳

	//使用time.Unix()可以将时间戳转换为时间格式
	//第一个参数是秒数，第二个参数是纳秒数
	timeObj := time.Unix(now.Unix(), 0)
	fmt.Println(timeObj)
	fmt.Println(timeObj.Date())
	fmt.Println(timeObj.Year())
	fmt.Println(timeObj.Month())
	fmt.Println(timeObj.Day())
	fmt.Println(timeObj.Hour())
	fmt.Println(timeObj.Minute())
	fmt.Println(timeObj.Second())

	//时间间隔  Duration类型代表两个时间点之间经过的时间，以纳秒为单位
	//在源码中定义为常量类型
	//可表示的最长字段大约为290年
	//表示一秒钟
	fmt.Println(time.Second)
	//在现在的时间上 加24小时
	fmt.Println(now.Add(24 * time.Hour))

	//定时器 - 每隔多少时间会将当前的时间返回
	//timer := time.Tick(time.Second) //每隔一秒钟，就将当前的时间返回
	//for t := range timer {
	//	fmt.Println(t)
	//}

	//时间格式化
	//自带的Format方法
	//Go语言中格式化时间模版不是使用常见的 Y-m-d H:M:S
	//而是使用Go语言的诞生时间2006年1月2号15点04分05秒
	//也就是2006 1 2 3(下午) 4 5
	//格式化时间，把语言中的时间对象转换成字符串类型的时间
	//将当前时间格式化为 Y-m-d
	fmt.Println(now.Format("2006-01-02"))
	//将当前时间格式化为 Y/m/d H:M:S(如果不想保留两位，直接1这样的，那就不用加0)
	fmt.Println(now.Format("2006/1/2 15:4:5"))
	//如果想表示成12进制的小时，将15换成03，3的话输出不保留两位格式
	//后面加上AM/PM，表示上午/下午
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	//精确到毫秒，在秒后面直接加上  .000
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))

	//利用time.Parse()  按照对应的格式，解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02", "2024-05-18")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	//Sub 两个时间相减
	//两个参数都是时间对象，结果才是时间间隔
	//调用的是方法
	fmt.Println(now)
	fmt.Println(timeObj)
	tSub := now.Sub(timeObj)
	fmt.Println(tSub)

	//Sleep()
	//参数是时间间隔，程序暂停一个时间间隔
	fmt.Println("开始sleep了")
	time.Sleep(1 * time.Minute)
	fmt.Println("10秒钟过去了...")
	n := 10000000000
	//强转成时间间隔类型的数据
	//转出来的时间间隔是纳秒
	fmt.Println("开始sleep了")
	time.Sleep(time.Duration(n))
	fmt.Println("...10000000000纳秒过去了")
}
