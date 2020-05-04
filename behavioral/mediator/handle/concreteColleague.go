package handle

import (
	"fmt"
	"github.com/treeforest/go-patterns/behavioral/mediator"
)

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
