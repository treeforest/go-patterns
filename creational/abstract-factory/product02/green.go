package product02

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"fmt"
)

type green struct {
	afbase.Color
}

func (p *green) Fill() {
	fmt.Println("Inside Green::fill() method.")
}
