package product02

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"sync"
)

type product02Factory struct {
	afbase.AbstractFactory
}

var once sync.Once
var instance *product02Factory = nil

func CreateProduct02Factory() afbase.AbstractFactory {
	once.Do(func() {
		instance = new(product02Factory)
	})
	return instance
}

func (p *product02Factory) CreateColor() afbase.Color {
	return new(green)
}

func (p *product02Factory) CreateShape() afbase.Shape {
	return new(circle)
}
