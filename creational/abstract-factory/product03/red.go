package product03

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"fmt"
)

type red struct {
	afbase.Color
}

func (p *red) Fill() {
	fmt.Println("Inside Red::fill() method.")
}
