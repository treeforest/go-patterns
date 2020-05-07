package main

import (
	"fmt"
	"reflect"
)

type (
	// 抽象状态类，定义一个接口以封装与 Context 的一个特定状态相关的行为
	State interface {
		Handle(Context)
	}

	// 抽象 Context 类
	Context interface {
		GetState() State
		SetState(State)
		Request()
	}
)

type (
	// 具体状态类，每一个子类实现一个与 Context 的状态相关的行为
	ConcreteStateA struct {}
	ConcreteStateB struct {}
	ConcreteStateC struct {}

	// 具体的 Context 类，维护一个具体状态类的实例，这个实例定义当前的状态
	ConcreteContext struct {
		state State
	}
)

func (o *ConcreteStateA) Handle(ctx Context) {
	// 设置 ConcreteStateA 的下一状态是 ConcreteStateB
	ctx.SetState(new(ConcreteStateB))
}

func (o *ConcreteStateB) Handle(ctx Context) {
	// 设置 ConcreteStateB 的下一状态是 ConcreteStateC
	ctx.SetState(new(ConcreteStateC))
}

func (o *ConcreteStateC) Handle(ctx Context) {
	// 设置 ConcreteStateC 的下一状态是 ConcreteStateD
	ctx.SetState(new(ConcreteStateA))
}

func CreateContext(state State) Context {
	c := new(ConcreteContext)
	c.state = state
	fmt.Println("初始状态：", reflect.Indirect(reflect.ValueOf(state)).Type().Name())
	return c
}

func (o *ConcreteContext) GetState() State {
	// 读取当前状态
	return o.state
}

func (o *ConcreteContext) SetState(state State) {
	// 设置新状态
	o.state = state
	fmt.Println("当前状态：", reflect.Indirect(reflect.ValueOf(state)).Type().Name())
}

func (o *ConcreteContext) Request() {
	// 对请求做处理，并设置下一状态
	o.state.Handle(o)
}

func main() {
	c := CreateContext(new(ConcreteStateA))

	c.Request()
	c.Request()
	c.Request()
	c.Request()
}