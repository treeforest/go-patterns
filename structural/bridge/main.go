package main

import (
	"github.com/treeforest/go-patterns/structural/bridge/brand"
	"github.com/treeforest/go-patterns/structural/bridge/soft"
)

func main() {
	hw := brand.CreateHuaWei()

	hw.SetSoft(new(soft.Game))
	hw.Run()

	hw.SetSoft(new(soft.AddressList))
	hw.Run()

	oppo := brand.CreateOppo()

	oppo.SetSoft(new(soft.Game))
	oppo.Run()

	oppo.SetSoft(new(soft.AddressList))
	oppo.Run()
}
