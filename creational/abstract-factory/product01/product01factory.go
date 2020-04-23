package product01

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"sync"
)

type product01Factory struct {
	afbase.AbstractFactory
}

var once sync.Once
var instance *product01Factory = nil

func CreateProduct01Factory() afbase.AbstractFactory {
	once.Do(func() {
		instance = new(product01Factory)
	})
	return instance
}

func (p *product01Factory) CreateColor() afbase.Color {
	return new(blue)
}

func (p *product01Factory) CreateShape() afbase.Shape {
	return new(rectangle)
}
