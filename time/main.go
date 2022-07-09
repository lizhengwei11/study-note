package main

import (
	"fmt"
	"time"
)

//标准库time用法
func main() {
	//time.Now()获取当前时间
	t := time.Now()
	fmt.Println(t)
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	fmt.Println(year, month, day, hour, minute, second)
	// 时间戳  从1970.1.1到现在的一个数
	timeStamp := t.Unix()      //总秒数
	timeStamp2 := t.UnixNano() //总纳秒数
	fmt.Println(timeStamp, timeStamp2)

	// 将时间戳转换为具体的时间格式
	tt := time.Unix(1657292038, 0)
	fmt.Println(tt)

	// 时间间隔
	n := 5 // int  int64不能相乘 time.Duration是基于int64造的自己的类型
	fmt.Println(time.Duration(n) * time.Second)

	// add()  t 当前时间 Add()添加一个小时
	t1 := t.Add(time.Hour * 2)
	fmt.Println(t1)
	// Sub 时间相减
	fmt.Println(t1.Sub(t))
	// Equal 判断两个时间是否相等
	fmt.Println(t1.Equal(t))
	// Before 判断两个时间的前后  返回bool值
	fmt.Println(t1.Before(t)) //再之前
	// After 判断两个时间的前后  返回bool值
	fmt.Println(t1.After(t)) //在之后

	// 定时器
	for tmp := range time.Tick(time.Second * 3) {
		fmt.Println(tmp)
	}

	// 时间格式化 2006 01 02 03 15
	println(t.Format("2006.01.02"))
	println(t.Format("2006.01.02 15:04:05"))
	println(t.Format("2006.01.02 15:04:05 PM")) //PM区分上下午

	//解析字符串类型时间
	str := "2022.07.08 23:32:38"
	// 1.先拿到时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(loc)
	// 2.根据时区去解析一个字符串的时间
	timeObj, _ := time.Parse("2006.01.02 15:04:05", str)
	fmt.Println(timeObj)
	// 3.解析哪个时区的时间
	timeObj1, _ := time.ParseInLocation("2006.01.02 15:04:05", str, loc) //解析字符串时间类型需要注意上面字符串年月日用什么分割下面解析传入时间格式也用字符串中的类型分割
	fmt.Println(timeObj1)

}
