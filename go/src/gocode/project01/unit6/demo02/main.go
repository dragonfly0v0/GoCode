package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误")
		panic(err)
	} else {
		fmt.Println("除法执行成功")
		fmt.Println("正常执行下面的逻辑")
	}
}

func test() (err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err is \n", err)
		}
	}()
	num1 := 10
	num2 := 0

	if num2 == 0 {
		return errors.New("num2 can not be zero")
	} else {
		rel := num2 / num1
		fmt.Println(rel)
		return nil
	}
}
