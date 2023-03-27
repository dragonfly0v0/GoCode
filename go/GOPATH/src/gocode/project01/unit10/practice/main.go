package main

import (
	"fmt"
)

// 定义结构体
type Matrix struct {
	Length uint16
	Width  uint16
}

// 1.编写一个方法，提供参数m,n,打印一个m*n的矩形
func (m Matrix) plot(Length uint16, Width uint16) {
	for i := 0; i < int(Length); i++ {
		fmt.Printf("+\t")
	}
	fmt.Println()
	for i := 0; i < int(Width); i++ {
		fmt.Printf("+\t")
		for z := 0; z < (int(Length) - 2); z++ {
			fmt.Printf(" \t")
		}
		fmt.Printf("+\t")
		fmt.Println()
	}
	for i := 0; i < int(Length); i++ {
		fmt.Printf("+\t")
	}
	fmt.Println()
}

// 2.编写方法计算矩形面积并返回
func (m Matrix) calculate(Length uint16, Width uint16) uint16 {
	res := Length * Width
	return res
}

func main() {
	var m Matrix
	m.plot(4, 6)
}
