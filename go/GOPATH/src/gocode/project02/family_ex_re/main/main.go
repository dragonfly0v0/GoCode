/*
编写文件TestMyAccount.go完成基本功能，要求：
1.完成显示主菜单，且能退出
2.完成显示明细的功能
3.完成登记收入的功能
4.完成登记支出的功能
*/
/*
改进：
1.用户输入4退出时，提示“确定要退出吗? y/n”, 必须输入正确的y/n，否则循环输入直到输入y或n
2.当没有任何收支明细时，提示“当前没有收支明细...来一笔吧”
3.在支出时判断余额是否足够，并给出相应提示提示
4.将面向过程的代码修改为面向对象的方法，编写MyFamilyAccount.go, 并使用testMyFamilyAccount.go去测试
*/
package main

import (
	"fmt"
)

type FamilyAccount struct {
	//声明必需的字段
	//声明变量接收用户输入的选项
	key string

	//声明变量控制是否退出for循环
	loop bool

	//定义初始账户余额
	rest float64

	//每次收支的金额
	money float64

	//字符串记录收支详情
	detail string

	//每次收支说明
	notes string

	flag string

	tag bool
}

func (this *FamilyAccount) ShowDetail() {
	fmt.Println("-------------当前收支明细记录------------")
	fmt.Println(this.detail)
}

func (this *FamilyAccount) ReceiveMoney() {
	fmt.Println("本次收入金额: ")
	fmt.Scanln(&this.money)
	this.rest += this.money //修改账户余额
	fmt.Println("本次收入说明: ")
	fmt.Scanln(&this.notes)
	//将收入情况拼接到detail
	this.detail += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.rest, this.money, this.notes)
	this.tag = true
}

func (this *FamilyAccount) SpendMoney() {
	fmt.Println("本次支出: ")
	fmt.Scanln(&this.money)
	if this.money > this.rest {
		fmt.Println("余额不足")
	}

	this.rest -= this.money //修改账户余额
	fmt.Println("本次支出说明: ")
	fmt.Scanln(&this.notes)
	//将收入情况拼接到detail
	this.detail += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.rest, this.money, this.notes)
	this.tag = false
}

func (this *FamilyAccount) Exit() {
	fmt.Println("你确定要退出吗? y/n")
	fmt.Scanln(&this.flag)
	switch this.flag {
	case "y":
		this.loop = false
	case "n":
		break
	default:
		fmt.Println("输入既不是y也不是n")
	}
}

func (this *FamilyAccount) ShowMenu() {
	myfamily_account := NewFamilyAccount()

	//显示主菜单
	for {
		fmt.Println("\n---------家庭收支记账软件----------")
		fmt.Println("          1.收支明细")
		fmt.Println("          2.登记收入")
		fmt.Println("          3.登记支出")
		fmt.Println("          4.退出")
		fmt.Print("请选择(1-4): ")

		fmt.Scanln(&myfamily_account.key)
		switch myfamily_account.key {
		case "1":
			if !myfamily_account.tag {
				fmt.Println("当前没有收支明细，来一笔吧")
			}
			myfamily_account.ShowDetail()
		case "2":
			myfamily_account.ReceiveMoney()
		case "3":
			myfamily_account.SpendMoney()
		case "4":
			myfamily_account.Exit()
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !myfamily_account.loop {
			break
		}
	}
	fmt.Println("退出记账软件...")
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		//声明变量接收用户输入的选项
		key: "",

		//声明变量控制是否退出for循环
		loop: true,

		//定义初始账户余额
		rest: 10000.00,
		//每次收支的金额
		money: 0.00,
		//字符串记录收支详情
		detail: "收支\t账户余额\t收支金额\t说   明",
		//每次收支说明
		notes: "",

		flag: "",

		tag: false,
	}
}

func main() {
	my := NewFamilyAccount()
	my.ShowMenu()
}
