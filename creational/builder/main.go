package main

import "github.com/treeforest/go-patterns/creational/builder/car"

func main() {
	assembly := car.NewCarBuilder()

	familyCar := assembly.Wheels(car.SportWheels).TopSpeed(50 * car.MPH).Color(car.RedColor).Build()
	familyCar.Drive()
	familyCar.Stop()

	sportCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Color(car.GreenColor).Build()
	sportCar.Drive()
	sportCar.Stop()
}
