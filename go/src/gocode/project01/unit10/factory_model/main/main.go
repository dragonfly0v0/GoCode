package main

import (
	"fmt"
	"gocode/project01/unit10/factory_model/model"
)

func main() {
	//创建Student实例
	var stu = model.Student{
		Name:  "Helen",
		Score: "89.76",
	}
	fmt.Println(stu)
}
