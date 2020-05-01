package soft

import (
	"fmt"
	"github.com/treeforest/go-patterns/structural/bridge"
)

type addressList struct {
}

func NewAddressList() bridge.ISoft {
	return new(addressList)
}

func (al *addressList) Run() error {
	fmt.Println("Run the handset address list.")
	return nil
}
