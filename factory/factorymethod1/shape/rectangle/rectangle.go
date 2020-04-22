package rectangle

import (
	"github.com/treeforest/design-pattern/factory/factorymethod1/shape"
	"fmt"
)

type Rectangle struct{
	shape.Shape
}

func (p *Rectangle)Draw(){
	fmt.Println("Inside Rectangle::draw() method.")
}
