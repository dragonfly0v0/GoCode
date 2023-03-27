package main

import (
	"encoding/json"
	"fmt"
)

// 定义学生结构体
type Monster struct {
	Name  string `json:"name"` //`json:"name"`就是struct tag
	Age   uint16 `json:"age"`
	Skill string `json:"skill"`
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

}
