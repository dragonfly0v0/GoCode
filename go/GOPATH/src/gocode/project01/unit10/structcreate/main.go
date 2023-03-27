package main

import "fmt"

//定义学生结构体
type Stu struct {
	Name string
	age  uint8
	sex  string
}

func main() {
	//方式2
	stu := Stu{"Jin", 20, "男"}
	fmt.Println(stu)

	//或者是
	stu.Name = "Cheng"
	stu.age = 22
	stu.sex = "男"

	//方式3：直接返回指针
	var stu2 *Stu = new(Stu)
	//stu2是一个指针，因此给字段赋值的标准方式是：
	//底层会进行处理，去掉*
	(*stu2).Name = "Hehai"
	stu2.age = 100

	//方式4
	var stu3 *Stu = &Stu{"Fav", 20, "男"}

	//或者
	var stu4 *Stu = &Stu{}
	//因为stu4是一个指针，用标准的访问字段的方法
	//也可以直接写为
	stu4.age = 29

	fmt.Println(stu2, stu3, stu4)
}
