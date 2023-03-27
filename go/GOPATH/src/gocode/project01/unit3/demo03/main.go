package main

import "fmt"

func main() {
	//定义一个变量
	var age int = 18
	fmt.Println("age address is: ", &age)

	var ptr *int = &age
	fmt.Println("指针*ptr指向的数值为: \n", *ptr)
	fmt.Println("ptr指向的数值为: \n", ptr)

}
