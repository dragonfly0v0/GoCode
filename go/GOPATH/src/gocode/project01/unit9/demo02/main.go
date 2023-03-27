package main

import "fmt"

func main() {
	//使用for-range遍历map
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "苏州"

	for _, val := range cities {
		fmt.Printf("city is %s\n", val)
	}
}
