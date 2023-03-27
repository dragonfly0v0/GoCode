/*
项目采用分级菜单方式。主菜单如下：
--------客户信息管理---------
1 添加客户
2 修改客户
3 删除客户
4 客户列表
5 退出
请选择[1-5]:

请选择[1-5]:1
-------------添加客户-------------
姓名:张三
性别:男
年龄:30
电话:010-367846738
邮箱:zsan@abc.com
-------------修改完成-------------

请选择[1-5]:2
-------------修改客户-------------
姓名(张三):<直接回车表示不修改>
性别(男):
年龄(30):
电话(010-367846738):
邮箱(zhang@abc.com):zsan@abc.com
-------------修改完成-------------

请选择[1-5]:3
-------------删除客户-------------
请选择待删除客户编号(-1退出):1
确认是否删除(y/n): y
-------------删除完成-------------

请选择[1-5]:4
-------------客户列表-------------
编号    姓名  性别  年龄  电话                邮箱
1      张三   男    30   010-367846738     zsan@abc.com
2
3
---------------------------------
*/
package service

import (
	"fmt"
	"gocode/project02/client_info_management_sysytem/model"
)

// 完成对Customer操作，包括增删改查
type CustomerService struct {
	customers []model.Customer

	//声明一个字段表示当前切片含有多少客户
	//该字段还可以作为新客户的Id+1
	Num int
}

// 编写方法，返回CustomerService实例
func NewCustomerService() *CustomerService {
	cus_service := &CustomerService{}
	cus_service.Num = 1
	customer := model.NewCustomer(1, "张三", "男", 40, "112", "zs@sohu.com")
	cus_service.customers = append(cus_service.customers, *customer)
	return cus_service
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// this一定要使用引用类型，不然会丢掉
func (this *CustomerService) Add(customer model.Customer) bool {
	this.Num++
	customer.Id = this.Num
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Change(index int, customer model.Customer) bool {
	fmt.Println(index)
	this.customers[index] = customer
	return true
}

// 根据Id查找客户在切片中对应下标，没有返回-1
func (this *CustomerService) FindbyId(id int) int {
	index := -1
	//遍历切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (this *CustomerService) Del(id int) bool {
	index := this.FindbyId(id)

	if index == -1 {
		fmt.Println("没有该客户...")
		return false
	}

	//如何从切片中删除元素
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
