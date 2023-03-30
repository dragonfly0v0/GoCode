package main

import "fmt"

func main() {
	//定义map
	var a map[int]string
	//只声明map内存没有分配空间
	//必须通过make函数对其初始化，内存才会分配空间
	a = make(map[int]string, 10) //map可以存放10个键值对
	//将键值存入map中
	a[12] = "sadwe"
	a[14] = "alen"

	fmt.Println(a)
}
