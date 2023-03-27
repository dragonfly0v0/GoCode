package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) Getsum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct {
		fmt.Println("Expect struct")
		return
	}

	num := val.NumField()
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为%v\n", i, val.Field(i))

		// 获取struct标签，通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为%v\n", i, tagVal)
		}
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	// var params []reflect.Value
	val.Method(1).Call(nil)

	// 调用结构体的第1个方法, 反射使用的方法0和方法1排序使用的是对比结构体方法首字母是默认按照ASCII码比较
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) // 传入参数是 []reflect.Value
	fmt.Println("res=", res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "黄精怪",
		Age:   800,
		Score: 90.35,
	}
	TestStruct(a)
}
