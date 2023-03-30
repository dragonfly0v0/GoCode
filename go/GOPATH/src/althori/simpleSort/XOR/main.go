package main

import "fmt"

// 异或
// 一个数组中有两种数是奇数个，其他数都是偶数个
// 找到这两种数
/*
	思路：
	1.所有的数异或，假设要找的数为a和b，则最后的结果为eor = a ^ b
	2.假设a和b在第i位不同即第i位异或结果为1，整个数组可以分为2类：第8为为1的数和第8位为0的数
	3.对第八位为1的数进行异或运算，则eor2 = a | b
*/
func PrintOdd_TimesNUm2(slice []int) {
	eor := 0
	for i := 1; i < len(slice); i++ {
		eor ^= slice[i] // 将两种奇数a^b得到
	}

	// eor = a ^ b
	// eor != 0
	// eor 必然在某位上是1
	rightOne := eor & (^eor + 1) // ^err按位取反, 取出最右侧的1
	// eor:    101011100
	// ^eor:   010100011
	// ^eor+1: 010100100
	// eor & (^eor + 1): 000000100   补码

	onlyOne := 0 // eor2
	for i := 1; i < len(slice); i++ {
		if (slice[i] & rightOne) == 0 { //最右侧不为1的取出来
			onlyOne ^= slice[i] // 0^slice[i]=slice[i]
		}
	}
	TheOtherOdd := eor ^ onlyOne
	fmt.Printf("一类是%v, 一类是%v\n", onlyOne, TheOtherOdd)

}

func main() {

}
