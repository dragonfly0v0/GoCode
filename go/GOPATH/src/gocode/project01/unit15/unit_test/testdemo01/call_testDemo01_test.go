package cal

import (
	"testing"
	//"gocode/project01/unit15/unit_test/testDemo01"
)

// 编写测试用例测试addUpper()
func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("addUpper(%v)执行错误")
		t.Fatalf("addUpper(%v)执行错误, 期望值55, 输出值为%v\n", 10, res)
	}

	// 如果正确, 记录日志
	t.Logf("执行正确\n")
}
