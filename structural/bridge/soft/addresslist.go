package soft

import (
	"fmt"
	"github.com/treeforest/go-patterns/structural/bridge"
)

// 通讯录
type addressList struct {
}

func NewAddressList() bridge.Soft {
	return new(addressList)
}

func (al *addressList) Run() error {
	fmt.Println("Run the handset address list.")
	return nil
}
