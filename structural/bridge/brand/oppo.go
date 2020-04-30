package brand

import (
	"github.com/treeforest/go-patterns/structural/bridge"
)

type oppo struct {
	soft bridge.ISoft
}

func CreateOppo() bridge.IBrand {
	return new(oppo)
}

func (op *oppo) SetSoft(soft bridge.ISoft) {
	// do something
	op.soft = soft
}

func (op *oppo) Run() {
	// do something
	op.soft.Run()
}