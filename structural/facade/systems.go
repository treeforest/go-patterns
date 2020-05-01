package main

import "fmt"

type SubSystemOne struct {
	//...
}

func (one *SubSystemOne) MethodOne() {
	fmt.Println("子系统方法一")
}

type SubSystemTwo struct {
	//...
}

func (two *SubSystemTwo) MethodTwo() {
	fmt.Println("子系统方法二")
}

type SubSystemThree struct {
	//...
}

func (three *SubSystemThree) MethodThree() {
	fmt.Println("子系统方法三")
}

type SubSystemFour struct {
	//...
}

func (four *SubSystemFour) MethodFour() {
	fmt.Println("子系统方法四")
}
