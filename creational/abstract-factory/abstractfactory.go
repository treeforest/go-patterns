package factory

// abstract factory
type Factory interface {
	CreateColor() Color
	CreateShape() Shape
}

type Color interface {
	Fill()
}

// 接口
type Shape interface {
	Draw()
}
