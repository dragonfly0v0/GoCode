package model

import "fmt"

type person struct {
	Name string
	age  uint8
	sal  float64
}

// 创建工厂模式函数
func NewPerson(name string) *person {
	return &person{ // &person{}，此时是一个指向person类型结构体实例的指针。
		Name: name,
	}
}

// 访问age和sal，编写Set和Get方法
func (p *person) SetAge(age uint8) {
	if age > 0 && age < 100 {
		p.age = age
	} else {
		fmt.Println("年龄不正确...")
	}
}

func (p *person) GetAge() uint8 {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal > 0 {
		p.sal = sal
	} else {
		fmt.Println("工资不是负数...")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
