package main

import (
	"fmt"
	"reflect"
)

var Student struct {
	Name  string
	Age   int
	Score float64
}

func test01(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println("Type is ", t)

	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println("kind=", k)
	fmt.Printf("v is %v, v's type is %T\n", v, v)
	fmt.Println("value=", v.Int())

	// 将v转换为interface{}
	iv := v.Interface()

	// 将interface{}断言转换为需要的类型
	num := iv.(int)
	fmt.Println(num)
}

func main() {
	test01(100)

	// stu  := Student{
	// 	Name: "tom",
	// 	Age: 20,
	// 	Score: 77.9,
	// }
}
