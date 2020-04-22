package product03

import (
	"fmt"
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
)

type red struct {
	handle.Color
}

func (p *red) Fill() {
	fmt.Println("Inside Red::fill() method.")
}

