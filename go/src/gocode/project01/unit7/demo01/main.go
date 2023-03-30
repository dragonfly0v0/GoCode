package main

import "fmt"

func main() {
	//给出5个学生分数，求出总和，并给出平均数
	//定义一个数组
	var score [5]int
	var sum int
	var avg int

	for i := 0; i < len(score); i++ {
		sum += score[i]
	}
	avg = sum / len(score)
	fmt.Println(avg)
}
