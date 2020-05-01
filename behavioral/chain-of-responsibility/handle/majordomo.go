package handle

import (
	"github.com/treeforest/go-patterns/behavioral/chain-of-responsibility"
	"fmt"
)

// 总监
type majordomo struct {
	manager
}

func CreateMajordomo(name string) handler.Manager {
	m := new(majordomo)
	m.name = name
	return m
}

func (m *majordomo) RequestApplications(request handler.Request) {
	if(request.GetRequestType() == "请假" && request.GetRequestNumber() <= 5) {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	}else {
		if m.superior != nil {
			// 将申请转交到上级
			m.superior.RequestApplications(request)
		}
	}
}