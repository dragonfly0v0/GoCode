package main

import "fmt"

func Sum(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	rel := Sum(20, 30)

	fmt.Println("rel is \n", rel)
}
