# Bridge Pattern

合成/聚合复用原则（CARP）：尽量使用合成/聚合，尽量不要使用类继承。  
  
桥接模式：将抽象部分和它的实现部分分离，使他们都可以独立地变化(实现指的是抽象类和它的派生类用来实现自己的对象)。

# Implementation

Abstraction 将 client 的请求转发给它的 Implementor 对象 

## Abstraction

```
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
```

## Implementor

```
// 实现类的接口，该接口不一定要与 Abstraction 一致，事实上两个接口可以完全不同。
// 一般来讲，Implementor 接口仅提供基本操作，而 Abstraction 则定义了基于这些操作的较高层次的操作
type Implementor interface {
	OperationImpl()
}

// ConcreteImplementorA 与 ConcreteImplementorB 实现 Implementor 接口并定义它的具体实现
type ConcreteImplementorA struct {}

func (o *ConcreteImplementorA) OperationImpl() {
	fmt.Println("ConcreteImplementorA OperationImpl")
}

type ConcreteImplementorB struct {}

func (o *ConcreteImplementorB) OperationImpl() {
	fmt.Println("ConcreteImplementorB OperationImpl")
}
```

# Usage

```
var abstraction Abstraction

abstraction = &RefinedAbstraction{
    imp: new(ConcreteImplementorA),
}
abstraction.Operation()

abstraction = &RefinedAbstraction{
    imp: new(ConcreteImplementorB),
}
abstraction.Operation()
```

```
ConcreteImplementorA OperationImpl
ConcreteImplementorB OperationImpl
```

# Rules of Thumb

桥接模式理解：实现系统可能有多角度分类，每一种分类有可能变化，那么就把这种多角度分离出来让它们独立变化，减少他们的耦合。
也就是，Abstraction 与 Implementor下面可以类似树形往下扩展，而两边互不影响（只要都实现该"树根"Abstraction 与 Implementor的接口）