# Abstract Factory Pattern

抽象工程模式：提供一个创建一系列相关或相互依赖对象的接口，而无需指定他们具体的类。

# Implementation

## Types

```
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
```

## Different Implementations

```
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
```

```
type rectangle struct {
}

func (p *rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}
```

```
type blue struct {
}

func (p *blue) Fill() {
	fmt.Println("Inside Blue::fill() method.")
}
```

* product02与product03实现类似，详见源码

# Usage

```
fmt.Println("product01: ")
prd01 := product01.CreateProduct01Factory()
c01 := prd01.CreateColor()
c01.Fill()
s01 := prd01.CreateShape()
s01.Draw()

fmt.Println("\nproduct02: ")
prd02 := product02.CreateProduct02Factory()
c02 := prd02.CreateColor()
c02.Fill()
s02 := prd02.CreateShape()
s02.Draw()

fmt.Println("\nproduct03: ")
prd03 := product03.CreateProduct03Factory()
c03 := prd03.CreateColor()
c03.Fill()
s03 := prd03.CreateShape()
s03.Draw()
```

# Rules of Thumb 

抽象工厂有利于产品族的扩展，不利于单个产品族内的产品扩展。