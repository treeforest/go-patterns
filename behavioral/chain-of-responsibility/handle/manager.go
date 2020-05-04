package handle

import "github.com/treeforest/go-patterns/behavioral/chain-of-responsibility"

// 管理者超类
type manager struct {
	handler.Manager
	superior handler.Manager
	name     string
}

func (m *manager) SetSuperior(superior handler.Manager) {
	m.superior = superior
}
