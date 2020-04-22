package product01

import (
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
	"fmt"
)

type rectangle struct {
	handle.Shape
}

func (p *rectangle)Draw(){
	fmt.Println("Inside Rectangle::draw() method.")
}
