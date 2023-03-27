package main

import (
	"fmt"
)

func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool, 值是%v\n", i, x)
		case float64, float32:
			fmt.Printf("param #%d is a float64, 值是%v\n", i, x)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is a uint, 值是%v\n", i, x)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is a int, 值是%v\n", i, x)
		case string:
			fmt.Printf("param #%d is a string, 值是%v\n", i, x)
		case nil:
			fmt.Printf("param #%d is a nil, 值是%v\n", i, x)
		default:
			fmt.Printf("param #%d is unknown, 值是%v\n", i, x)
		}
	}
}

func main() {
	var n1 float32 = 90.32
	var str string = "sdhi"
	var slice []int
	addr := "beijing"

	TypeJudge(n1, str, slice, addr)

}
