package main

import "fmt"

func Sum(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func main() {
	rel := Sum
	fmt.Printf("rel type is: %T, Sum type is: %T \n", rel, Sum)
}
