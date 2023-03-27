package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat1 := Cat{Name: "tom", Age: 18}
	cat2 := Cat{Name: "bob", Age: 8}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10

	//取出
	cat := <-allChan
	// cat的类型未知，需要使用断言
	// fmt.Println(cat.Name) //如果未使用断言，cat.Name undefined (type interface{} has no field or method Name)
	a := cat.(Cat)
	fmt.Println(a.Name)
}
