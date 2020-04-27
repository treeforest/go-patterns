package main

import (
	"fmt"
	"sync"
)

// 若要使用代理功能，则必须实现该接口的相同方法
type IObject interface {
	ObjDo(action string)
}

// 实际进行数据处理的类
type Object struct {
	IObject
	//...
}

// ObjDo是实现于IObject的接口
// 实现了真正的逻辑处理
func (obj *Object) ObjDo(action string) {
	fmt.Printf("I can, %s", action)
}

// 具备有拦截行为的代理类
type ProxyObject struct {
	IObject
	once sync.Once
	object *Object
}

// ObjDo是实现于IObject的接口
// 目的是为了拦截访问真正的Object的行为
func (p *ProxyObject) ObjDo(action string) {
	p.once.Do(func() {
		p.object = new(Object)
	})

	// do something

	p.object.ObjDo(action)
}
