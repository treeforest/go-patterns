package main

import "fmt"

// 外观类
type facade struct {
	one *SubSystemOne
	two *SubSystemTwo
	three *SubSystemThree
	four *SubSystemFour
}

func CreateFacade() *facade {
	f := new(facade)
	f.one = new(SubSystemOne)
	f.two = new(SubSystemTwo)
	f.three = new(SubSystemThree)
	f.four = new(SubSystemFour)
	return f
}

func (p *facade) MethodA() {
	fmt.Println("方法组 MethodA()")
	p.one.MethodOne()
	p.two.MethodTwo()
	p.four.MethodFour()
}

func (p *facade) MethodB() {
	fmt.Println("方法组 MethodB()")
	p.two.MethodTwo()
	p.three.MethodThree()
}