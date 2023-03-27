package model

// 定义结构体
type Student struct {
	Name  string
	Score float64
}

type stu struct {
	Name string
	sex  string
}

// 通过工厂模式解决首字母小写无法被其他包调用的问题
func Newstu(n string, s string) *stu {
	return &stu{
		Name: n,
		sex:  s,
	}
}

// 首字母小写不能访问，使用方法设置成可以访问
func (s *stu) GetSex() string {
	return s.sex
}
