package main

import (
	"fmt"
	)

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

func main() {
	o := CreateObjectStructure()
	o.Attach(new(ConcreteElementA))
	o.Attach(new(ConcreteElementB))

	v1 := new(ConcreteVisitor1)
	v2 := new(ConcreteVisitor2)

	o.Accept(v1)
	o.Accept(v2)
}
