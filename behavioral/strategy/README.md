# Strategy Pattern

策略模式：它定义了算法家族，分别封装起来，让它们之间可以互相替换，此模式让算法的变化，不会影响到使用算法的客户。

# Implementation

```
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

```

# Usage

```
o := new(Operation)

o.SetOperator(new(Addition))
fmt.Println("5 + 3 =", o.Operate(5,3))

o.SetOperator(new(Subtraction))
fmt.Println("5 - 3 =", o.Operate(5,3))

o.SetOperator(new(Multiplication))
fmt.Println("5 * 3 =", o.Operate(5, 3))
```

# Rule of Thumb

策略模式是一种定义一系列算法的方法，从概念上来看，所有这些算法完成的都是相同的工作，只是实现不同，它可以以相同的方式调用所有的算法，减少了各种算法类与使用算法类之间的耦合。

策略模式的 Operator 类层次为 Operation 定义了一系列的可供重用的算法或行为，继承有助于析取出这些算法中的公共功能。

策略模式的有点是简化了单元测试，因为每个算法都有自己的类，可以通过自己的接口单独测试。