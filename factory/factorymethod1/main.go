package main

import "github.com/treeforest/design-pattern/factory/factorymethod1/shapefactory"

func main() {
	sf := shapefactory.CreateShapeFactory()

	circle := sf.CreateCircle()
	circle.Draw()

	rectangle := sf.CreateRectangle()
	rectangle.Draw()

	square := sf.CreateSquare()
	square.Draw()
}
