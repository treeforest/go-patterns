package product03

import (
	"github.com/treeforest/go-patterns/creational/abstract-factory/afbase"
	"sync"
)

type product03Factory struct {
	afbase.AbstractFactory
}

var once sync.Once
var instance *product03Factory = nil

func CreateProduct03Factory() afbase.AbstractFactory {
	once.Do(func() {
		instance = new(product03Factory)
	})
	return instance
}

func (p *product03Factory) CreateColor() afbase.Color {
	return new(red)
}

func (p *product03Factory) CreateShape() afbase.Shape {
	return new(square)
}
