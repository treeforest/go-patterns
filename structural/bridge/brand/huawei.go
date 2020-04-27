package brand

import (
	"github.com/treeforest/go-patterns/structural/bridge/soft"
	)

type huawei struct {
	IBrand
	soft soft.ISoft
}

func CreateHuaWei() IBrand {
	return new(huawei)
}

func (hw *huawei) SetSoft(soft soft.ISoft) {
	// do something
	hw.soft = soft
}

func (hw *huawei) Run() {
	// do something
	hw.soft.Run()
}