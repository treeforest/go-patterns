package handle

type AbstractFactory interface {
	CreateColor() Color
	CreateShape() Shape
}