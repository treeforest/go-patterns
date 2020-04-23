package afbase

type AbstractFactory interface {
	CreateColor() Color
	CreateShape() Shape
}
