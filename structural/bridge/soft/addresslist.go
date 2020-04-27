package soft

import "fmt"

type AddressList struct {
	ISoft
}

func (al *AddressList) Run() error {
	fmt.Println("Run the handset address list.")
	return nil
}