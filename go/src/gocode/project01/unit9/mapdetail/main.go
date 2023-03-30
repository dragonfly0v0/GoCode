package main

//定义学生结构体
type Stu struct {
	Name string
	age  uint8
	sex  string
}

func main() {
	students := make(map[int]Stu, 10)

	//创建两个学生
	stu1 := Stu{
		Name: "Su",
		age:  18,
		sex:  "男",
	}

	stu2 := Stu{
		Name: "Huan",
		age:  26,
		sex:  "女",
	}

	students[0] = stu1
	students[1] = stu2
}
