package main

import (
	"fmt"
)

type AintFace interface {
	Say()
}

type Bintface interface {
	Hello()
}

type Student struct {
	Name string
}

func (stu Student) Say() {
	fmt.Println("stu Say")
}

type integer int

type Monster struct {
}

func (i integer) Say() {
	fmt.Println("interger say i = ", i)
}

func (m Monster) Say() {
	fmt.Println(m, "Say")
}

func (m Monster) Hello() {
	fmt.Println("Hello")
}

func main() {
	var stu Student //结构体变量实现了Say() 实现了AintFace
	var a AintFace = stu
	a.Say()

	var i integer = 10
	var b AintFace = i
	b.Say()

	//Monster实现了Aintface和Bintface
	var monster Monster
	var A AintFace = monster
	var B Bintface = monster
	A.Say()
	B.Hello()

}
