package main

// 定义抽象类的接口
type Abstraction interface {
	Operation()
}

// 维护一个指向 Implementor 类型对象的指针，并实现 Abstraction 接口
type RefinedAbstraction struct {
	imp Implementor
}

func (o *RefinedAbstraction) Operation() {
	o.imp.OperationImpl()
}
