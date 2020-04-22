package square

import (
	"github.com/treeforest/design-pattern/factory/factorymethod1/shape"
	"fmt"
)

type Square struct{
	shape.Shape
}

func (p *Square) Draw(){
	fmt.Println("Inside Square::draw() method.")
}