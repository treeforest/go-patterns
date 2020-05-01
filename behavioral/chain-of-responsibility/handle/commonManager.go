package handle

import (
	"github.com/treeforest/go-patterns/behavioral/chain-of-responsibility"
	"fmt"
)

// 经理
type commonManager struct {
	manager
}

func CreateCommonManager(name string) handler.Manager {
	m := new(commonManager)
	m.name = name
	return m
}

func (m *commonManager) RequestApplications(request handler.Request) {
	if(request.GetRequestType() == "请假" && request.GetRequestNumber() <= 2) {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	}else {
		if m.superior != nil {
			// 将申请转交到上级
			m.superior.RequestApplications(request)
		}
	}
}