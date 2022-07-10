package main

// Date 后端使用的是int64  进行json转换的时候与反序列·化的时候使用字符串
type Date struct {
	ID int64 `json:"id,string"`
}
