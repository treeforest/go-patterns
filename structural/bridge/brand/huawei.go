package brand

import (
	"github.com/treeforest/go-patterns/structural/bridge"
)

type huawei struct {
	soft bridge.ISoft
}

func CreateHuaWei() bridge.IBrand {
	return new(huawei)
}

func (hw *huawei) SetSoft(soft bridge.ISoft) {
	// do something
	hw.soft = soft
}

func (hw *huawei) Run() {
	// do something
	hw.soft.Run()
}