package store

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
1.编写一个Monster结构体，字段Name, Age, Skill
2.给Monster绑定方法Store，可以将一个Monster变量序列化后保存到文件
3.给Monster绑定方法ReStore，可以将一个序列化的Monster，从文件中读取，并反序列化为Monster对象
4.编写测试文件store_test.go, 编写测试函数TestStore和TestRestore进行测试
*/
type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (moster Monster) Store(m *Monster) {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("序列化失败 err=%v\n", err)
	}
	fmt.Printf("序列化结果是%v\n", string(data))

	file := "F:/code/store.txt"
	info, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Opne file err: ", err)
		return
	}
	// 及时关闭file句柄
	defer info.Close()

	writer := bufio.NewWriter(info)
	writer.WriteString(string(data))
	writer.Flush()
}

func (moster Monster) Restore() bool {
	file := "F:/code/store.txt"
	info, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Opne file err: ", err)
		return false
	}

	var monster Monster
	err3 := json.Unmarshal([]byte(info), &monster)
	if err3 != nil {
		fmt.Printf("反序列化失败 err=%v\n", err3)
		return false
	}
	fmt.Printf("反序列化后的结果是%v\n", monster)
	return true
}
