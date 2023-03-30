package main

import "fmt"

func main() {
	test()
	fmt.Println("除法执行成功")
	fmt.Println("正常执行下面的逻辑")
}

func test() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err is \n", err)
		}
	}()
	num1 := 0
	num2 := 10
	rel := num2 / num1
	fmt.Println(rel)

}
