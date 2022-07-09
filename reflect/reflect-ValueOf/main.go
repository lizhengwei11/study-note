package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)
	k := v.Kind() //拿到值对应类型的种类
	//如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		//把反射取到的值转换成一个float32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v,%T\n", ret, ret)
	case reflect.Int32:
		//把反射取到的值转换成一个int32类型的变量
		ret := int32(v.Int())
		fmt.Printf("%v,%T\n", ret, ret)
	}

}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)
	k := v.Elem().Kind() //Elem根据指针取对应的值
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}

}
func main() {
	//var aa int32 = 100
	//reflectValue(aa)
	//var bb float32 = 12.001
	//reflectValue(bb)
	var cc int32 = 10
	reflectSetValue(&cc)
	var a *int
	reflect.ValueOf(a).IsNil() //.IsNil()  用来判断指针是否为空
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("var a *int nil 返回值是否有效：:", reflect.ValueOf(nil).IsValid()) // IsValid判断返回值是否有效

	b := struct{}{}
	reflect.ValueOf(b).FieldByName("aaa").IsValid()  //FieldByName().IsValid().  用来查找结构体里的字段
	reflect.ValueOf(b).MethodByName("ccc").IsValid() //MethodByName().IsValid().  用来查找结构体里的方法
	fmt.Println("b := struct{}{} IsValid:", reflect.ValueOf(a).IsNil())
	c := map[string]int{}
	fmt.Println("c := map[string]int{} 不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("xx")).IsValid())
}
