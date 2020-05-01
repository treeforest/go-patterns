package main

import (
	"github.com/treeforest/go-patterns/creational/builder"
	"github.com/treeforest/go-patterns/creational/builder/car"
)

func main() {
	assembly := car.NewCarBuilder()

	familyCar := assembly.Wheels(builder.SportWheels).TopSpeed(50 * builder.MPH).Color(builder.RedColor).Build()
	familyCar.Drive()
	familyCar.Stop()

	sportCar := assembly.Wheels(builder.SteelWheels).TopSpeed(150 * builder.MPH).Color(builder.GreenColor).Build()
	sportCar.Drive()
	sportCar.Stop()
}
