package main

import (
	"fmt"
)

// 编写一个学生考试系统
type Student struct {
	Name  string
	Age   uint8
	Score float64
}

//将pupil和graduate共有方法绑定到* Student

func (stu *Student) ShowInfo() {
	fmt.Printf("姓名: %v, 年龄: %v, 分数: %v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64) {
	stu.Score = score
}

func (stu *Student) ShowScore() float64 {
	return stu.Score
}

type Pupil struct {
	Student //嵌入Student匿名结构体
}

func (stu *Pupil) testing() {
	fmt.Println("小学生正在考试中....")
}

type Graduate struct {
	Student
}

func (stu *Graduate) testing() {
	fmt.Println("大学生正在考试中....")
}

func main() {
	pupil := &Pupil{
		Student{
			Name: "Tom",
			Age:  10,
		},
	}

	pupil.testing()
	pupil.Student.SetScore(90.88)
	pupil.Student.ShowInfo()

}
