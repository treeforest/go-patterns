package brand

import (
	"github.com/treeforest/go-patterns/structural/bridge/soft"
	)

type oppo struct {
	IBrand
	soft soft.ISoft
}

func CreateOppo() IBrand {
	return new(oppo)
}

func (op *oppo) SetSoft(soft soft.ISoft) {
	// do something
	op.soft = soft
}

func (op *oppo) Run() {
	// do something
	op.soft.Run()
}