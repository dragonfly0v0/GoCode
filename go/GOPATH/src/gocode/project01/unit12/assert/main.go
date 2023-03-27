package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point

	var b Point
	b = a.(Point)
	fmt.Println(b)

	var t float32
	var x interface{}
	x = t
	//转成float，带检查
	_, ok := x.(float32)
	if ok == true {
		fmt.Print("convert success")
	} else {
		fmt.Println("convert fail")
	}
}
