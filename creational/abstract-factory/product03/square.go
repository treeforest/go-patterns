package product03

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"fmt"
)

type square struct {
	afbase.Shape
}

func (p *square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}
