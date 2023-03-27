package main

import (
	"fmt"
)

func BubbleSort(array []int) {
	var temp int

	for i := 0; i < (len(array) - 1); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] <= array[j] {
				continue
			} else {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
		fmt.Printf("第%d轮排序后的结果是%d\n", i, array)
	}
	fmt.Println(array)
}

// func main() {
// 	var array [10]int
// 	rand.Seed(time.Now().Unix())
// 	for i := 0; i < len(array); i++ {
// 		array[i] = rand.Intn(1000) //0 <= n < 1000
// 	}
// 	fmt.Println(array)

// 	BubbleSort(array[:])

// }
