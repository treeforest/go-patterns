package product02

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory"
	"sync"
)

type product02Factory struct {

}

var once sync.Once
var instance *product02Factory = nil

func CreateProduct02Factory() factory.Factory {
	once.Do(func() {
		instance = new(product02Factory)
	})
	return instance
}

func (p *product02Factory) CreateColor() factory.Color {
	return new(green)
}

func (p *product02Factory) CreateShape() factory.Shape {
	return new(circle)
}
