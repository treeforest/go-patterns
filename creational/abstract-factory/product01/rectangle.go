package product01

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"fmt"
)

type rectangle struct {
	afbase.Shape
}

func (p *rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}
