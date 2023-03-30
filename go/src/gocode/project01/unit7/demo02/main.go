package main

import "fmt"

func main() {
	//给出5个学生分数，求出总和，并给出平均数
	//定义一个数组
	var score [5]float32
	var sum float32
	var avg float32

	for i := 0; i < len(score); i++ {
		fmt.Printf("请录入第%d个学生成绩: \n", i)
		fmt.Scanln(&score[i])
	}

	for key, val := range score {
		fmt.Printf("第%d个学生的成绩为%2f", key, val)
	}
	avg = sum / float32(len(score))
	fmt.Println(avg)
}
