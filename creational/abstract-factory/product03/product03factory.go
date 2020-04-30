package product03

import (
	"sync"
	"github.com/treeforest/go-patterns/creational/abstract-factory"
)

type product03Factory struct {

}

var once sync.Once
var instance *product03Factory = nil

func CreateProduct03Factory() factory.Factory {
	once.Do(func() {
		instance = new(product03Factory)
	})
	return instance
}

func (p *product03Factory) CreateColor() factory.Color {
	return new(red)
}

func (p *product03Factory) CreateShape() factory.Shape {
	return new(square)
}
