package brand

import (
	"github.com/treeforest/go-patterns/structural/bridge"
)

type huawei struct {
	soft bridge.Soft
}

func CreateHuaWei() bridge.Brand {
	return new(huawei)
}

func (hw *huawei) SetSoft(soft bridge.Soft) {
	// do something
	hw.soft = soft
}

func (hw *huawei) Run() {
	// do something
	hw.soft.Run()
}
