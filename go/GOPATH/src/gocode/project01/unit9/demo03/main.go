package main

import "fmt"

func main() {
	//使用map记录monster的信息name和age，一个monster对应一个map，且monster的个数可以动态的增加 ==> map切片
	//map切片的使用

	//1.声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2) //放2个monster

	//2.增加妖怪信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "东皇太一"
		monsters[0]["race"] = "妖"
	}
	fmt.Println(monsters)

	//如果增加多个monster，则需要切片的append函数，动态增加monster
	//1.先定义monster信息
	newMonster := map[string]string{
		"name": "孙悟空",
		"race": "神",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)
}
