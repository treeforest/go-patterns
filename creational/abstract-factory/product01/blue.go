package product01

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"fmt"
)

type blue struct {
	afbase.Color
}

func (p *blue) Fill() {
	fmt.Println("Inside Blue::fill() method.")
}
