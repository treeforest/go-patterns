package product02

import (
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
	"sync"
)

type product02Factory struct {
	handle.AbstractFactory
}

var once sync.Once
var instance *product02Factory = nil
func CreateProduct02Factory() handle.AbstractFactory{
	once.Do(func() {
		instance = new(product02Factory)
	})
	return instance
}

func (p *product02Factory) CreateColor() handle.Color {
	return new(green)
}

func (p *product02Factory) CreateShape() handle.Shape {
	return new(circle)
}