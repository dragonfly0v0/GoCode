package main

import "fmt"

func main() {
	var age int //声明int
	age = 18
	fmt.Println("age = ", age)

	//声明和赋值可以合成一句
	var size int = 100
	fmt.Println("size is ", size)

	fmt.Println("--------------------------------------")

	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	var n4, name, n5 = 10, "Alen", 8.9
	fmt.Println(n4, name, n5)
}
