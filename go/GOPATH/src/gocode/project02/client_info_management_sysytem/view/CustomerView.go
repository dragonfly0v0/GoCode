package main

import (
	"fmt"
	"gocode/project02/client_info_management_sysytem/model"
	"gocode/project02/client_info_management_sysytem/service"
)

type CustomerView struct {
	key  string
	loop bool
	flag string

	//增加CustomerService字段
	customerService *service.CustomerService
}

// 显示所有客户信息
func (this *CustomerView) List() {
	customers := this.customerService.List()

	fmt.Println("-------------客户列表-------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}

	fmt.Println("---------------------------------")
}

// 添加客户信息
func (this *CustomerView) AddCustomer() {

	fmt.Println("-------------添加客户-------------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email := ""
	fmt.Scanln(&email)

	//构建新的Customer实例
	//注意：id号系统分配，没有让用户输入
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if this.customerService.Add(customer) {
		fmt.Println("-------------添加完成-------------")

	} else {
		fmt.Println("-------------添加失败-------------")

	}

}

// 修改客户信息
func (this *CustomerView) UpdateCustomer() {

	fmt.Println("-------------修改客户-------------")
	this.List()

	id := -1
	fmt.Scanln(&id)

	index := this.customerService.FindbyId(id)
	if index == -1 {
		fmt.Println("该客户不存在")
		return
	} else {
		customers := this.customerService.List()
		name := ""
		fmt.Printf("姓名(%v): ", customers[index].Name)
		fmt.Scanln(&name)

		gender := ""
		fmt.Printf("性别(%v): ", customers[index].Gender)
		fmt.Scanln(&gender)

		age := 0
		fmt.Printf("年龄(%v): ", customers[index].Age)
		fmt.Scanln(&age)

		phone := ""
		fmt.Printf("电话(%v): ", customers[index].Phone)
		fmt.Scanln(&phone)

		email := ""
		fmt.Printf("邮箱(%v): ", customers[index].Email)
		fmt.Scanln(&email)

		customer := model.NewCustomer(customers[index].Id, name, gender, age, phone, email)
		fmt.Printf("====%T, %v====", customer, customer)
		if this.customerService.Change(index, *customer) {
			fmt.Println("-------------修改成功-------------")

		} else {
			fmt.Println("-------------修改失败-------------")

		}
	}
}

// 删除客户信息
func (this *CustomerView) DelCustomer() {
	fmt.Println("-------------删除客户-------------")
	this.List()

	num := -1
	fmt.Println("请选择待删除客户编号(-1退出):")
	fmt.Scanln(&num)

	if num == -1 {
		return
	}

	fmt.Println("你确定要删除吗? y/n")
	fmt.Scanln(&this.flag)
	switch this.flag {
	case "y":
		if this.customerService.Del(num) {
			fmt.Println("-------------删除成功-------------")

		} else {
			fmt.Println("-------------删除失败-------------")

		}
	case "n":
		break
	default:
		fmt.Println("输入既不是y也不是n")
	}

}

// 显示主菜单
func (this *CustomerView) MainMenu() {
	for {
		fmt.Println("--------客户信息管理---------")
		fmt.Println("1 添加客户")
		fmt.Println("2 修改客户")
		fmt.Println("3 删除客户")
		fmt.Println("4 客户列表")
		fmt.Println("5 退出")
		fmt.Println("请选择[1-5]:")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.AddCustomer()
		case "2":
			this.UpdateCustomer()
		case "3":
			this.DelCustomer()
		case "4":
			this.List()
		case "5":
			this.loop = false
		default:
			fmt.Println("输入超出范围, 请重新输入...")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("退出客户管理系统...")
}

func main() {
	//创建CustomerView实例
	csview := CustomerView{
		key:  "",
		loop: true,
		flag: "",
	}

	//初始化csview的customerService字段
	csview.customerService = service.NewCustomerService()

	//显示所有客户信息

	//显示主菜单
	csview.MainMenu()
}
