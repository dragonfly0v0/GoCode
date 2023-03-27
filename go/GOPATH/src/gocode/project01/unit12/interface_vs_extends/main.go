package main

import (
	"fmt"
)

// 猴子不仅能继承老猴子的技能
// 还可以通过接口实现鸟的飞行和鱼的游泳
// 定义Monkey struct
type Monkey struct {
	Name string
}

func (m *Monkey) climb() {
	fmt.Println(m.Name, "生来会爬树")
}

type Bird interface {
	fly()
}

type Fish interface {
	Swim()
}

func (m *Monkey) fly() {
	fmt.Println(m.Name, "学习后会飞")
}

func (m *Monkey) Swim() {
	fmt.Println(m.Name, "学习后游泳")
}

// 定义little_monkey struct
type LittleMonkey struct {
	Monkey
}

func main() {
	//创建Little_Monkey实例
	var monkey LittleMonkey
	monkey.Name = "小猴子"
	monkey.climb()
	monkey.fly()

}
