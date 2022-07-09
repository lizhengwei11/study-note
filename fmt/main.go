package main

import "fmt"

func main() {
	//输出字符串
	fmt.Print("输出字符串")
	fmt.Println("换行输出字符串")
	//格式化输出 %s 字符串 %d整型 %t布尔类型 \n换行
	fmt.Printf("年龄%d姓名%s\n是人吗\n%t", 10, "李四", true)
	// %f 浮点类型占位符 %6.1f 指定宽度为6精度为1
	fmt.Printf(" %6.1f", 12.111)

	//输入，获取用户输入的值
	var name string

	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Printf("%s", name)
}
