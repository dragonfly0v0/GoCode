package main

import (
	"fmt"
	"sort"
)

func main() {
	//map排序
	map1 := make(map[int]string, 10)
	map1[0] = "唐僧"
	map1[1] = "孙悟空"
	map1[2] = "猪八戒"
	map1[3] = "沙僧"
	map1[4] = "白龙马"
	map1[5] = "金翅大鹏"
	map1[20] = "玉兔"
	map1[15] = "锦毛鼠"

	fmt.Println(map1)
	//按照map的key顺序排序输出
	//1.先将map的key放入切片
	//2.对切片排序
	//3.遍历切片，按照key输出map的value
	var keys []int
	for key, _ := range map1 {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(map1[k])
	}
}
