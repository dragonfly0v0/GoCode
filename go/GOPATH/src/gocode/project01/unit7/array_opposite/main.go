package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//随机生成5个数，将其反转打印
	//1.随机生成5个数，rand.Intn()
	//2.放到数组
	//3.反转打印
	var array [10]int

	//为了每次生成的随机数不同，都需要一个seed值
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(1000) //0 <= n < 1000
	}
	fmt.Println(array)

	for i := 0; i < len(array)/2; i++ {
		temp := array[len(array)-1-i]
		array[len(array)-1-i] = array[i]
		array[i] = temp
	}
	fmt.Println(array)
}
