package main

import "fmt"

// 实现类的接口，该接口不一定要与 Abstraction 一致，事实上两个接口可以完全不同。
// 一般来讲，Implementor 接口仅提供基本操作，而 Abstraction 则定义了基于这些操作的较高层次的操作
type Implementor interface {
	OperationImpl()
}

// ConcreteImplementorA 与 ConcreteImplementorB 实现 Implementor 接口并定义它的具体实现
type ConcreteImplementorA struct{}

func (o *ConcreteImplementorA) OperationImpl() {
	fmt.Println("ConcreteImplementorA OperationImpl")
}

type ConcreteImplementorB struct{}

func (o *ConcreteImplementorB) OperationImpl() {
	fmt.Println("ConcreteImplementorB OperationImpl")
}
