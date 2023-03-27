package main

import "reflect"

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)

	// 改变rVal的值
	rVal.Elem().SetInt(20)
}
