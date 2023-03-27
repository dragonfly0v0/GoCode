package main

import (
	"encoding/json"
	"fmt"
)

// 演示将结构体、map和切片序列化
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Birth string
	Sal   string
	Skill string
}

// 将json字符串反序列化成struct
func unMarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":5000,\"Birth\":\"十万年前\",\"Sal\":\"5000\",\"Skill\":\"法天象地\"}"
	// 真实情况下str是由网络传输得到

	// 声明monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化失败 err=%v\n", err)
	}
	fmt.Printf("反序列化后的结果是%v\n", monster)
}

// 将json字符串反序列化成map
func unMarshalMap() {
	str := "{\"addr\":\"火云洞\",\"age\":2000,\"name\":\"红孩儿\"}"

	var a map[string]interface{}
	a = make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("反序列化失败 err=%v\n", err)
	}
	fmt.Printf("反序列化后的结果是%v\n", a)
}

func unMarshalSlice() {
	str := "[{\"addr\":\"花果山\",\"age\":5000,\"name\":\"孙悟空\"},{\"addr\":[\"云栈洞\",\"高老庄\"],\"age\":5200,\"name\":\"猪八戒\"}]"

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("反序列化失败 err=%v\n", err)
	}
	fmt.Printf("反序列化后的结果是%v\n", slice)
}

func main() {
	unMarshalStruct()
	unMarshalMap()
	unMarshalSlice()
}
