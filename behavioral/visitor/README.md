# Visitor Pattern

访问者模式：表示一个作用于某一对象结构中的各元素的操作。它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。

# Implementation

```$xslt
type (
	Visitor interface {
		VisitConcreteElementA(element Element)
		VisitConcreteElementB(element Element)
	}

	Element interface {
		Accept(visitor Visitor)
	}
)

type (
	ConcreteVisitor1 struct {}
	ConcreteVisitor2 struct {}

	ConcreteElementA struct {}
	ConcreteElementB struct {}
)

func (v *ConcreteVisitor1) VisitConcreteElementA(e Element) {
	fmt.Println("ConcreteVisitor1 visits ConcreteElementA.")
}

func (v *ConcreteVisitor1) VisitConcreteElementB(e Element) {
	fmt.Println("ConcreteVisitor1 visit ConcreteElementB.")
}

func (v *ConcreteVisitor2) VisitConcreteElementA(e Element) {
	fmt.Println("ConcreteVisitor2 visits ConcreteElementA.")
}

func (v *ConcreteVisitor2) VisitConcreteElementB(e Element) {
	fmt.Println("ConcreteVisitor2 visit ConcreteElementB.")
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

type ObjectStructure struct {
	mp map[Element]struct{}
}

func CreateObjectStructure() ObjectStructure {
	return ObjectStructure{
		mp: make(map[Element]struct{}),
	}
}

func (o *ObjectStructure) Attach(element Element) {
	o.mp[element] = struct{}{}
}

func (o *ObjectStructure) Detach(element Element) {
	delete(o.mp, element)
}

func (o *ObjectStructure) Accept(visitor Visitor) {
	for e, _ := range o.mp {
		(e.(Element)).Accept(visitor)
	}
}

```

# Usage

```$xslt
o := CreateObjectStructure()
o.Attach(new(ConcreteElementA))
o.Attach(new(ConcreteElementB))

v1 := new(ConcreteVisitor1)
v2 := new(ConcreteVisitor2)

o.Accept(v1)
o.Accept(v2)
```

```$xslt
ConcreteVisitor1 visits ConcreteElementA.
ConcreteVisitor1 visit ConcreteElementB.
ConcreteVisitor2 visits ConcreteElementA.
ConcreteVisitor2 visit ConcreteElementB.
```

# Rule of Thumb

访问者模式的目的是要把处理从数据结构分离出来，很多系统可以按照算法和数据结构分开，如果这样的系统有比较稳定的数据结构，又有易于变化的算法的话，使用访问者模式就是比较合适的，因为访问者模式使得算法操作的增加变得容易。

访问者模式的优点：增加新的操作很容易，因为增加新的操作就意味着增加一个新的访问者，访问者模式将有关的行为集中到一个访问者对象中。

访问者模式的缺点：使增加新的数据结构边困难。