package main

import "fmt"

func main() {
	//数组的初始化方式
	//第一种：
	var arr1 [5]int = [5]int{72, 34, 29}
	fmt.Println(arr1)

	//第二种：
	var arr2 = [5]int{72, 34, 29}
	fmt.Println(arr2)

	//第三种：
	var arr3 = [...]int{72, 34, 29} //长度不确定
	fmt.Println(arr3)

	//第四种：
	var arr4 = [...]int{2: 66, 9: 20, 12: 40, 13: 20} //下标2的位置是66
	fmt.Println(arr4)
}
