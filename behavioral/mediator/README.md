# Mediator Pattern

中介者模式(调停者模式):用一个中介对象来封装一系列的对象交互。中介者使各对象不需要显示地相互作用，从而使其耦合松散，而且可以独立地改变它们之间的交互。

# Implementation

## Types

```$xslt
package mediator

// 抽象中介者接口
type Mediator interface {
	Send(msg string, c Colleague) error
}

// 抽象同事接口
type Colleague interface {
	Send(msg string) error
	Notify(msg string) error
}
```

## Different Implementations

```$xslt
// 具体中介者类
type concreteMediator struct {
	C1 mediator.Colleague
	C2 mediator.Colleague
}

func CreateConcreteMediator() *concreteMediator {
	return new(concreteMediator)
}

func (cm *concreteMediator) Send(msg string, c mediator.Colleague) error {
	// 中介者实现消息如何转发的逻辑
	if c == cm.C1 {
		cm.C2.Notify(msg)
	} else {
		cm.C1.Notify(msg)
	}
	return nil
}
```

```$xslt
type concreteColleague1 struct {
	m mediator.Mediator
}

func CreateConcreteColleague1(m mediator.Mediator) mediator.Colleague {
	cc := new(concreteColleague1)
	cc.m = m
	return cc
}

func (cc *concreteColleague1) Send(msg string) error {
	// 同事类对象间的消息都是通过中介者转发
	cc.m.Send(msg, cc)
	return nil
}

func (cc *concreteColleague1) Notify(msg string) error {
	fmt.Println("同事 1 得到消息: ", msg)
	return nil
}
```

```
type concreteColleague2 struct {
	m mediator.Mediator
}

func CreateConcreteColleague2(m mediator.Mediator) mediator.Colleague {
	cc := new(concreteColleague2)
	cc.m = m
	return cc
}

func (cc *concreteColleague2) Send(msg string) error {
	cc.m.Send(msg, cc)
	return nil
}

func (cc *concreteColleague2) Notify(msg string) error {
	fmt.Println("同事 2 得到消息: ", msg)
	return nil
}
```

# Usage

```$xslt
m := handle.CreateConcreteMediator()

c1 := handle.CreateConcreteColleague1(m)
c2 := handle.CreateConcreteColleague2(m)

m.C1 = c1
m.C2 = c2

c1.Send("吃过饭了吗？")
c2.Send("没有呢，你打算请客？")
```

```$xslt
同事 2 得到消息:  吃过饭了吗？
同事 1 得到消息:  没有呢，你打算请客？
```

# Rule of Thumb

中介者模式很容易在系统中应用，也很容易在系统中误用。当系统出现了“多对多”交互复杂的对象群时，不要着急于使用中介者模式，而要先反思你的系统在设计上是不是合理。

优点：Mediator的出现减少了各个Colleague的耦合性，使得可以独立地改变和复用各个Colleague类的Mediator;其次，由于把对象如何协作进行了抽象，将中介作为一个独立的概念并将其封装在一个对象中，这样关注的对象就从对象各自本身的行为转移到它们之间的交互上来，也就是站在一个更宏观的角度去看待系统。

缺点：由于concreteMediator控制集中化，于是就把交互复杂性变为了中介者的复杂性，这就使得中介者会变得比任何一个concreteColleague都复杂