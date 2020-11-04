package main

import "fmt"

// 抽象组件接口，方便动态地添加职责
type Component interface {
	Operation()
}

// 具体组件对象，给这个对象添加一些职责
type ConcreteComponent struct {}
func (cc *ConcreteComponent) Operation() {
	fmt.Println("ConcreteComponent Operation")
}

// 维持一个指向 Component 对象的指针，并定义一个与 Component 接口一致的接口
type Decorator struct {
	Component Component
}

func (o *Decorator) SetComponent(component Component) {
	o.Component = component
}

func (o *Decorator) Operation() {
	o.Component.Operation()
}

// 向组件添加职责
type ConcreteDecoratorA struct {
	Dec Decorator
}

func (o *ConcreteDecoratorA) Operation() {
	fmt.Println("ConcreteDecoratorA Operation")
	o.Dec.Operation()
}

type ConcreteDecoratorB struct {
	Dec Decorator
}

func (o *ConcreteDecoratorB) Operation() {
	fmt.Println("ConcreteDecoratorB Operation")
	o.Dec.Operation()
}