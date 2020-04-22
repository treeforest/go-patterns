package product03

import (
	"fmt"
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
)

type square struct{
	handle.Shape
}

func (p *square) Draw(){
	fmt.Println("Inside Square::draw() method.")
}