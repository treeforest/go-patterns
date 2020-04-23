package product02

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"fmt"
)

type circle struct {
	afbase.Shape
}

func (p *circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}
