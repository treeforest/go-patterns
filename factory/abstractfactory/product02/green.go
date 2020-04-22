package product02

import (
	"fmt"
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
)

type green struct {
	handle.Color
}

func (p *green) Fill() {
	fmt.Println("Inside Green::fill() method.")
}