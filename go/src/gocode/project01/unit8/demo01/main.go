package main

import "fmt"

func main() {
	//二维数组
	var arr = [6]int{1, 3, 5, 8, 9, 10}

	//切片构建在数组之上
	//定义一个切片slice，[]动态变化的数组长度不写，int类型，arr是原数组
	var slice []int = arr[2:6]
	fmt.Println(slice)
	//切片长度

	//make内置函数创建切片。
	//基本语法：var切片名[type = make([], len, [cap])]
	slice1 := make([]int, 4, 20)
	fmt.Println(len(slice1))

	//定义一个切片，直接指定具体数组，使用类似make的方式

	//切片容量
	fmt.Println(cap(slice))
}
