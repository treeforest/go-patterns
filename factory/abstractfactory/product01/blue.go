package product01

import (
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
	"fmt"
)

type blue struct {
	handle.Color
}

func (p *blue) Fill() {
	fmt.Println("Inside Blue::fill() method.")
}
