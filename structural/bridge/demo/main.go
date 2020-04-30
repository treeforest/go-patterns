package main

import (
	"github.com/treeforest/go-patterns/structural/bridge/brand"
	"github.com/treeforest/go-patterns/structural/bridge/soft"
)

func main() {
	hw := brand.CreateHuaWei()

	hw.SetSoft(soft.NewGame())
	hw.Run()

	hw.SetSoft(soft.NewAddressList())
	hw.Run()

	oppo := brand.CreateOppo()

	oppo.SetSoft(soft.NewGame())
	oppo.Run()

	oppo.SetSoft(soft.NewAddressList())
	oppo.Run()
}
