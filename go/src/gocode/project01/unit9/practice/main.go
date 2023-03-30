package main

import "fmt"

// 定义学生结构体
type Stu struct {
	Name string
	age  uint8
	sex  string
}

func modifyUser(users map[string]map[string]string, name string) {
	defalt_passwd := "888"

	if _, ok := users[name]; ok {
		if defalt_passwd != users[name]["passwd"] {
			users[name]["passwd"] = defalt_passwd
		}
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["passwd"] = defalt_passwd
		users[name]["nickname"] = "昵称~" + name

	}
	fmt.Println(users)

}

func main() {
	user := make(map[string]map[string]string)

	modifyUser(user, "Kehan")
}
