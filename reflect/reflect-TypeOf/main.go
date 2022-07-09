package main

import (
	"fmt"
	"reflect"
)

// reflect TypeOf
/**
.Kind()  李文周博客有详解
*/
func reflectType(x interface{}) {
	//不知道别人调用我这个函数会传进来什么值
	//方式一： 类型断言 (猜类型)
	//方式二： 借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind()) //Name 具体信息 Kind种类信息
	fmt.Printf("%T\n", obj)
}

type cat struct {
}
type dog struct {
}

func main() {
	var a float64 = 1.333333
	reflectType(a)
	var b int16 = 10
	reflectType(b)
	//结构体类型
	var c cat
	var d dog
	reflectType(c)
	reflectType(d)
	var e []string
	var f []int
	reflectType(e) //切片没有  Go语言反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
	reflectType(f)

}
