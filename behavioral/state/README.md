# State Pattern

状态模式：当一个对象的内在状态改变时允许改变其行为，这个对象看起来像是改变了其类。

# Implementation

```$xslt
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
```

# Usage

```$xslt
c := CreateContext(new(ConcreteStateA))

c.Request()
c.Request()
c.Request()
c.Request()
```

```$xslt
初始状态： ConcreteStateA
当前状态： ConcreteStateB
当前状态： ConcreteStateC
当前状态： ConcreteStateA
当前状态： ConcreteStateB
```

# Rule of Thumb

状态模式主要解决的是当控制一个对象状态转换的条件表达式过于复杂的情况。把状态的判断逻辑转移到表示不同状态的一系列类中，可以把复杂的判断逻辑简化。当然，如果这个状态判断很简单，那就没必要用“状态模式”了。

状态模式的好处是将于特定状态相关的行为局部化，并且将不同状态的行为分割开来。用于消除庞大的条件分支语句。

状态模式通过把各种状态逻辑分布到 State 的子类之间，来减少相互间的依赖。

当一个对象的行为取决于它的状态，并且它必须在运行时刻根据状态改变它的行为时，就可以考虑使用状态模式了。