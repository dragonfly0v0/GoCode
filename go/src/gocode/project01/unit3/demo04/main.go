package main

import "fmt"

func main() {
	//实现连加

	sum := 0

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}

	fmt.Println("sum is \n", sum)
}
