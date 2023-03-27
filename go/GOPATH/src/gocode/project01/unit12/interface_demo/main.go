package main

import (
	"fmt"
)

// 声明一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type MobilePhone struct {
}

type Camera struct {
}

type Computer struct {
}

// MobilePhone实现Usb接口的方法
func (p MobilePhone) Start() {
	fmt.Println("手机启动...")
}

func (p MobilePhone) Stop() {
	fmt.Println("手机关闭...")
}

// Camera实现Usb接口的方法
func (p Camera) Start() {
	fmt.Println("相机启动...")
}

func (p Camera) Stop() {
	fmt.Println("相机关闭...")
}

// 给Computer编写一个方法working
// 所谓实现Usb接口，指实现Usb接口声明的所有方法
func (c Computer) Working(usb Usb) {
	//通过usb接口变量调用start和stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := MobilePhone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
