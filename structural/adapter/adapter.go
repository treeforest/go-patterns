package main

import "fmt"

// 客户端所期待的接口，目标可以是具体的或抽象的类，也可以是接口（这里示范的是接口）
type Target interface {
	Request()
}

// 适配器：通过在内部封装一个 Adaptee 对象，把源接口装换成目标接口
type Adapter struct {
	Adp *Adaptee
}

func (o *Adapter) Request() {
	o.Adp.SpecificRequest()
}

// 需要适配的类，原接口 SpecificRequest 已无法满足现有接口 Request 的需求
type Adaptee struct {}

func (o *Adaptee) SpecificRequest() {
	fmt.Println("Hello, This is SpecificRequest.")
}
