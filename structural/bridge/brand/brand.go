package brand

import "github.com/treeforest/go-patterns/structural/bridge/soft"

type IBrand interface {
	SetSoft(soft soft.ISoft)
	Run()
}
