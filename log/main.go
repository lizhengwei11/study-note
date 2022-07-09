package main

import (
	"log"
	"os"
)

//go标准库
//log 内置简单的日志库
func main() {

	//日志的配置
	log.SetPrefix("Text:")                                  //设置前缀
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds) //设置格式
	f, _ := os.OpenFile("./log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	log.SetOutput(f)

	log.Println("hello")

	//基本用法  内置log打印都会换行
	//log.Print("helo")
	//log.Printf("hello %s", "printf") //格式化打印
	//log.Println("hello ln")          //打印并换行

	for i := 1; i < 10; i++ {

		////log.Fatal  用法  打印日志后并退出
		//log.Fatal("helo")                //打印日志并退出
		//log.Fatalf("hello %s", "printf") //格式化打印
		//log.Fatalln("hello ln")

		// log.Panic 用法  打印日志详细信息并退出
		//log.Panic("hello")
		//log.Panicf("hello %s", "printf") //格式化打印
		//log.Panicln("hello ln")
	}
}
