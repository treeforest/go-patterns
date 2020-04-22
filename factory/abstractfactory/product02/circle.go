package product02

import (
	"fmt"
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
)

type circle struct{
	handle.Shape
}

func (p *circle) Draw(){
	fmt.Println("Inside Circle::draw() method.")
}
