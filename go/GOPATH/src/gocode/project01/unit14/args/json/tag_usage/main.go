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

func testStruct() {
	monster := Monster{
		Name:  "牛魔王",
		Age:   5000,
		Birth: "十万年前",
		Sal:   "5000",
		Skill: "法天象地",
	}

	// 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败 err=%v\n", err)
	}
	fmt.Printf("序列化结果是%v\n", string(data))
}

// 序列化map
func testMap() {
	var a map[string]interface{}
	// 使用map前需要先make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 2000
	a["addr"] = "火云洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化失败 err=%v\n", err)
	}
	fmt.Printf("序列化结果是%v\n", string(data))
}

// 切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	// 使用map前先make
	m1 = make(map[string]interface{})
	m1["name"] = "孙悟空"
	m1["age"] = 5000
	m1["addr"] = "花果山"
	//slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "猪八戒"
	m2["age"] = 5200
	m2["addr"] = [2]string{"云栈洞", "高老庄"}

	slice = append(slice, m1, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化失败 err=%v\n", err)
	}
	fmt.Printf("序列化结果是%v\n", string(data))
}

// 基本数据类型序列化
func testFloat64() {
	var num1 float64 = 98.237937
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化失败 err=%v\n", err)
	}
	fmt.Printf("序列化结果是%v\n", string(data))
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
