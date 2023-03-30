package main

import "fmt"

func main() {
	//二维数组
	var arr [2][3]int = [2][3]int{{1, 3, 5}, {8, 9, 10}}
	for key, val := range arr {
		fmt.Printf("------------------\n")
		fmt.Println(key, val)
		for key2, val2 := range val {
			fmt.Println(key2, val2)
			fmt.Printf("++++++++++++++++++\n")
		}
	}
}
