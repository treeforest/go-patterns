package handle

import (
	"fmt"
	"github.com/treeforest/go-patterns/behavioral/chain-of-responsibility"
)

type generalManager struct {
	manager
}

func CreateGeneralManager(name string) handler.Manager {
	m := new(generalManager)
	m.name = name
	return m
}

func (m *generalManager) RequestApplications(request handler.Request) {
	if request.GetRequestType() == "请假" {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	} else if request.GetRequestType() == "加薪" && request.GetRequestNumber() <= 500 {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	} else if request.GetRequestType() == "加薪" && request.GetRequestNumber() > 500 {
		fmt.Printf("%s:%s 数量%d 再说吧。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	}
}
