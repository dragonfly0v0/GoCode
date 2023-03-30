package main

import (
	"fmt"
	"math/rand"
	"time"
)

func QuickSort(slice []int, left, right int) {
	if left < right {
		privot := partion(slice, left, right)
		fmt.Printf("****循环1****\n")
		QuickSort(slice, left, privot-1)
		fmt.Printf("****循环1结束, 循环2开始****\n")
		QuickSort(slice, privot+1, right)
		fmt.Printf("****循环2结束****\n")
	}
}

func partion(slice []int, left, right int) int {
	pivot := slice[right]
	i := left - 1
	for j := left; j < right; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
		fmt.Printf("第%d次循环输出的slice是%d\n", j, slice)
	}
	slice[i+1], slice[right] = slice[right], slice[i+1]
	fmt.Printf("返回Quicksort时的slice是%d\n, return的值是%d", slice, i+1)
	fmt.Printf("-------------------------\n")
	return i + 1
}

func main() {
	var array [4]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000) //0 <= n < 1000
	}
	fmt.Println(array)

	QuickSort(array[:], 0, len(array)-1)
}
