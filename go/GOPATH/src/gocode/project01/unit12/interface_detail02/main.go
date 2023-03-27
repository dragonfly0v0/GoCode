package main

import "fmt"

type AintFace interface {
	test01()
}

type Bintface interface {
	test01()
}

type Cintface interface {
	AintFace
	Bintface
	test03()
}

// 如果实现C接口，需要将A,B接口的方法都实现
type Stu struct{}

func (stu Stu) test01() {
	fmt.Println("test01")
}

// func (stu Stu) test02() {
// 	fmt.Println("test02")
// }

func (stu Stu) test03() {
	fmt.Println("test03")
}

func main() {
	var stu Stu
	var a Cintface = stu
	a.test01()
	// a.test02()
	// a.test03()
}
