package main

import (
	"github.com/treeforest/design-pattern/factory/factorymethod2/circlefactory"
	"github.com/treeforest/design-pattern/factory/factorymethod2/rectanglefactory"
	"github.com/treeforest/design-pattern/factory/factorymethod2/squarefactory"
)

func main(){
	cf := new(circlefactory.CircleFactory)
	c := cf.CreateCircle()
	c.Draw()

	rf := new(rectanglefactory.RectangleFactory)
	r := rf.CreateRectangle()
	r.Draw()

	sf := new(squarefactory.SquareFactory)
	s := sf.CreateSquare()
	s.Draw()
}
