package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// json 序列化
/**
Marshal  序列化
Unmarshal 反序列化
*/

func main() {
	openBool()
	opInterface()
	opNull()
	opStruct()
	opStructUn()
	opEnCoder()
}
func openBool() {
	var ok bool
	ok = true
	// Marshal 序列化
	result, _ := json.Marshal(ok) //返回[]byte,err
	fmt.Println(string(result))

}

func opInterface() {
	var number interface{} //nil 按go语言没有null格式的 null是json转换的json格式的null
	result, _ := json.Marshal(number)
	fmt.Println(string(result))
}

func opNull() {
	var number interface{}
	_ = json.Unmarshal([]byte(`null`), &number)
	fmt.Println(number)
}

type Info struct { // 序列化时候如果name首字母为小写,因go的可见性，其他包读取不到，即json序列化和反序列化时,json读取不到
	Name  string  `json:"name"`       // json序列化后重命名
	Age   int     `json:"age,string"` // json序列化后转换为string类型
	Price float64 `json:"-"`          //- 不序列化操作
}

func opStruct() {
	var info Info
	info = Info{
		Name:  "张帅",
		Age:   44,
		Price: 12.1111,
	}
	result, _ := json.Marshal(info)
	fmt.Println(string(result))
}
func opStructUn() {
	var value interface{}
	_ = json.Unmarshal([]byte(`{"name":"张帅","age":"44"}`), &value)
	// json反序列转换为map
	fmt.Println(value)
	// Valid检查是否是标准的json数据
	if ok := json.Valid([]byte(`{"name":"zhang"}`)); !ok {
		fmt.Println("invalid")
	}
}

// opEnCoder 自定义json序列化规则
func opEnCoder() {
	_ = json.NewEncoder(os.Stdout).Encode(
		[]*Info{
			{"zz", 112, 122},
			{"sc", 123, 124},
			{"ss", 126, 127},
		},
	)
}
