package main

import (
	"fmt"
	"reflect"
)

type Stu struct {
	Name string
	Age  int
}

func ReflectStruct(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println("Type is ", t)
	fmt.Println("Kind is ", t.Kind())

	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println("kind=", k)
	fmt.Printf("v is %v, v's type is %T\n", v, v)

	// 将v转换为interface{}
	iv := v.Interface()

	// 将interface{}断言转换为需要的类型
	student, ok := iv.(Stu)
	if !ok {
		fmt.Println("error... ")
	}
	fmt.Println(student)
}

func main() {
	// 定义结构体实例
	stu := Stu{
		Name: "Tom",
		Age:  20,
	}

	ReflectStruct(stu)

}
