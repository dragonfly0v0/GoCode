package model

import "fmt"

//声明一个Customer struct, 表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//编写工厂模式，返回初始化结构体
func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) *Customer {

	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户信息
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", this.Id, this.Name, this.Gender,
		this.Age, this.Phone, this.Email)
	return info
}

//不带id
func NewCustomer2(name string, gender string, age int,
	phone string, email string) Customer {

	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
