package main

import (
	"encoding/json"
	"fmt"
)

// 定义结构体
type Monster struct {
	Name  string `json:"name"` //`json:"name"`就是struct tag
	Age   uint16 `json:"age"`
	Skill string `json:"skill"`
}

// 给Monster结构体添加speak方法
func (m Monster) speak() {
	fmt.Printf("抓住唐僧\n")
}

// 给Monster结构体添加calculate方法
func (m Monster) calculate(start int, end int) int {
	res := 0
	res = (start + end) * ((end - start + 1) / 2)

	fmt.Println("计算结果是: ", res)
	return res
}

func main() {
	var monster Monster
	monster.Name = "红孩儿"
	monster.Age = 800
	monster.Skill = "三昧真火"

	//将其进行json格式的序列号处理
	//json.Marshal使用到反射
	data, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json encoding err\n", err)
		return
	}
	fmt.Println("Json后的数据是: ", string(data)) //Json后的数据是:  {"Name":"红孩儿"}, N是大写的

	monster.speak()

	monster.calculate(1, 50)
}
