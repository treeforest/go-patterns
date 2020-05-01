package product01

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory"
	"sync"
)

type product01Factory struct {
}

var once sync.Once
var instance *product01Factory = nil

func CreateProduct01Factory() factory.Factory {
	once.Do(func() {
		instance = new(product01Factory)
	})
	return instance
}

func (p *product01Factory) CreateColor() factory.Color {
	return new(blue)
}

func (p *product01Factory) CreateShape() factory.Shape {
	return new(rectangle)
}
