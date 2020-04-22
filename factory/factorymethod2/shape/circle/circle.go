package circle

import (
	"github.com/treeforest/design-pattern/factory/factorymethod1/shape"
	"fmt"
)

type Circle struct{
	shape.Shape
}

func (p *Circle) Draw(){
	fmt.Println("Inside Circle::draw() method.")
}
