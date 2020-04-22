package product01

import (
	"github.com/treeforest/design-pattern/factory/abstractfactory/handle"
	"sync"
)

type product01Factory struct {
	handle.AbstractFactory
}

var once sync.Once
var instance *product01Factory = nil
func CreateProduct01Factory() handle.AbstractFactory{
	once.Do(func() {
		instance = new(product01Factory)
	})
	return instance
}

func (p *product01Factory) CreateColor() handle.Color{
	return new(blue)
}

func (p *product01Factory) CreateShape() handle.Shape {
	return new(rectangle)
}

