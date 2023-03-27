package main

import (
	"fmt"
)

func Date() {
	var year int
	var month int
	var day int
	for {
		fmt.Scan("请输入年份: ", &year)
		fmt.Scan("请输入月份: ", &month)
		fmt.Scan("请输入日: ", &day)

		if year < 0 {
			fmt.Print("年份输入错误")
		} else if month < 1 || month > 12 {
			fmt.Print("月份输入错误")
		} else if day < 1 || day > 31 {
			fmt.Print("日输入错误")
		}

		if year/400 == 0 && month == 2 {
			if day < 1 || day > 30 {
				fmt.Print("2月日期不能超过29号")
			}
		} else if month == 2 {
			if day < 1 || day > 29 {
				fmt.Print("2月日期不能超过28号")
			}
		}

		fmt.Printf("您输入的日期是: %d年%d月%d日\n", year, month, day)
	}
}

func main() {
	//循环打印输入的月份的天数。[使用continue实现]
	//要求有判断月份是否有错误的语句

}
