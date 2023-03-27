package main

import (
	"flag"
	"fmt"
)

// cmd>main.exe -f c:/aaa.txt -p 200 -u root

func main() {
	// 定义变量用于接收命令行参数值
	var user string
	var pwd string
	var host string
	var port int

	// &user就是接收命令行输入的-u后的参数值
	// "u" 接收-u 后面的指定参数
	// "" 默认值
	// "用户名, 默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名, 默认为空")
	flag.StringVar(&pwd, "p", "", "密码, 默认为空")
	flag.StringVar(&host, "h", "localhost", "用户名, 默认为localhost")
	flag.IntVar(&port, "port", 3306, "用户名, 默认为3306")

	// 重要操作, 转换, 必须调用该操作
	flag.Parse()

	// 输出结果
	fmt.Printf("user=%v, pwd=%v, host=%v, port=%v\n", user, pwd, host, port)
}
