package main

import "github.com/treeforest/go-patterns/behavioral/mediator/handle"

func main() {
	m := handle.CreateConcreteMediator()

	c1 := handle.CreateConcreteColleague1(m)
	c2 := handle.CreateConcreteColleague2(m)

	m.C1 = c1
	m.C2 = c2

	c1.Send("吃过饭了吗？")
	c2.Send("没有呢，你打算请客？")
}
