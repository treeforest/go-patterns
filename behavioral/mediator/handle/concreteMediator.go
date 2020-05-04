package handle

import "github.com/treeforest/go-patterns/behavioral/mediator"

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
