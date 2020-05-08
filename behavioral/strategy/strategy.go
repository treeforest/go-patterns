package main

import "fmt"

// 对整数操作的抽象接口（抽象策略接口）
type Operator interface {
	Apply(int, int) int
}

type (
	// 对整数操作的具体实现类。实现对整数操作的不同策略
	Addition struct {}
	Subtraction struct {}
	Multiplication struct {}
)

func (Addition) Apply(lVal, rVal int) int {
	return lVal + rVal
}

func (Subtraction) Apply(lVal, rVal int) int {
	return lVal - rVal
}

func (Multiplication) Apply(lVal, rVal int) int {
	return lVal * rVal
}

//
type Operation struct {
	op Operator
}

func (o *Operation) SetOperator(op Operator) {
	o.op = op
}

func (o *Operation) Operate(lVal, rVal int) int {
	return o.op.Apply(lVal, rVal)
}

func main() {
	o := new(Operation)

	o.SetOperator(new(Addition))
	fmt.Println("5 + 3 =", o.Operate(5,3))

	o.SetOperator(new(Subtraction))
	fmt.Println("5 - 3 =", o.Operate(5,3))

	o.SetOperator(new(Multiplication))
	fmt.Println("5 * 3 =", o.Operate(5, 3))
}