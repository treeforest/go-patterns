package product03

import (
		"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
	"sync"
)

type product03Factory struct {
	handle.AbstractFactory
}

var once sync.Once
var instance *product03Factory = nil
func CreateProduct03Factory() handle.AbstractFactory{
	once.Do(func() {
		instance = new(product03Factory)
	})
	return instance
}

func (p *product03Factory) CreateColor() handle.Color {
	return new(red)
}

func (p *product03Factory) CreateShape() handle.Shape {
	return new(square)
}
