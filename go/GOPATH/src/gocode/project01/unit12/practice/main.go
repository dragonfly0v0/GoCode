package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func bubble_sort(slice []int) []int {
	var temp int
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp = slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

type Hero struct {
	Name string
	age  int
}
type HeroSlice []Hero

// 实现接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法决定使用什么标准进行排序
// 标准是按照年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	if hs[i].age < hs[j].age {
		return true
	} else {
		return false
	}
}

func (hs HeroSlice) Swap(i, j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	//定义一个数组/切片
	//var arr [5]int
	var slice = []int{0, -20, -10, 2, 98}

	//对slice进行冒泡排序
	// bubble_sort(slice)
	// fmt.Println(slice)

	sort.IntSlice.Sort(slice)
	fmt.Println(slice)

	//对结构体切片进行排序
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("Hero~%d", rand.Intn(100)),
			age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	//排序前
	for _, val := range heros {
		fmt.Println(val)
	}

	//排序后
	sort.Sort(heros)
	fmt.Printf("-----------排序后-----------\n")
	//fmt.Println(heros)
	for _, val := range heros {
		fmt.Println(val)
	}
}
