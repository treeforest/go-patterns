package brand

import (
	"github.com/treeforest/go-patterns/structural/bridge"
)

type oppo struct {
	soft bridge.Soft
}

func CreateOppo() bridge.Brand {
	return new(oppo)
}

func (op *oppo) SetSoft(soft bridge.Soft) {
	// do something
	op.soft = soft
}

func (op *oppo) Run() {
	// do something
	op.soft.Run()
}
