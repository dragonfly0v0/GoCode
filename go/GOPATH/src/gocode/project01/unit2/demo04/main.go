package main

import "fmt"

func main() {
	//定义浮点类型数据
	var num1 float32 = 3.14
	fmt.Println(num1)

	var num2 float32 = -3.14
	fmt.Println(num2)

	var num3 float32 = 314e-2
	fmt.Println(num3)

	var num4 float32 = 314e+2
	fmt.Println(num4)

	var num5 float64 = 314e+2 //E大小写都可以
	fmt.Println(num5)
}
