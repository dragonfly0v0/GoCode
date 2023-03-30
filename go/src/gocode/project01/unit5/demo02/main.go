package main

import "fmt"

func Sum(x int, y int) (sum int, value int) {
	sum = x + y
	value = y - x
	return sum, value
}

// func Sum(x int, y int) (sum int) {
// 	sum = x + y
// 	value = y - x
// 	return sum
// }

func main() {
	rel, _ := Sum(20, 30)

	fmt.Println("rel is \n", rel)
}
