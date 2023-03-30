package main

import (
	"fmt"
	"strconv"
)

func main() {
	//golang的编码统一为utf-8,(ascii的字符(字母和数字)占1个字符, 汉字占3个字符)
	//统计字符长度
	str := "hello苏"
	fmt.Println(len(str))

	//字符串遍历
	//转换为rune的切片
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("字符=%c\n", str2[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("123")
	fmt.Printf("num is %d, err is %v\n", n, err)

	//整数转字符串
	str3 := strconv.Itoa(4783)
	fmt.Printf("%s, type is %T\n", str3, str3)

	//字符串转[]byte
	var bytes = []byte("hello go")
	fmt.Printf("%v\n", bytes)

	//[]byte转字符串
	str4 := string([]byte{99, 100})
	fmt.Printf("str4 is %s\n", str4)

}
