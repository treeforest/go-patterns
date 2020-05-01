# Builder Pattern

建造者模式：将一个复杂对象的构建和它的表示分离，使得相同的构建过程可以创建不同地表示。

# Implementation

示例实现了一辆车的建造过程。

## Types

```
package builder

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string

const (
	SportWheels Wheels = "sports"
	SteelWheels        = "steel"
)

type Builder interface {
	Color(c Color) Builder
	Wheels(w Wheels) Builder
	TopSpeed(s Speed) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

```

## Different Implementations

```
type car struct {
	color  builder.Color
	wheels builder.Wheels
	speed  builder.Speed
}

func (c *car) Drive() error {
	fmt.Printf("A %s car with %s tires is going at %f ", c.color, c.wheels, c.speed)
	if c.speed == builder.MPH {
		fmt.Println("MPH")
	} else {
		fmt.Println("KPH")
	}

	return nil
}

func (c *car) Stop() error {
	fmt.Println("The car stopped.")
	return nil
}
```

```
type carBuilder struct {
	c *car
}

func NewCarBuilder() builder.Builder {
	cb := new(carBuilder)
	cb.c = new(car)
	return cb
}

func (cb *carBuilder) Color(c builder.Color) builder.Builder {
	cb.c.color = c
	return cb
}

func (cb *carBuilder) Wheels(w builder.Wheels) builder.Builder {
	cb.c.wheels = w
	return cb
}

func (cb *carBuilder) TopSpeed(s builder.Speed) builder.Builder {
	cb.c.speed = s
	return cb
}

func (cb *carBuilder) Build() builder.Interface {
	return cb.c
}
```

# Usage

```
assembly := car.NewCarBuilder()

familyCar := assembly.Wheels(builder.SportWheels).TopSpeed(50 * builder.MPH).Color(builder.RedColor).Build()
familyCar.Drive()
familyCar.Stop()

sportCar := assembly.Wheels(builder.SteelWheels).TopSpeed(150 * builder.MPH).Color(builder.GreenColor).Build()
sportCar.Drive()
sportCar.Stop()
```

```
A red car with sports tires is going at 50.000000 KPH
The car stopped.
A green car with steel tires is going at 150.000000 KPH
The car stopped.
```

# Rules of Thumb

建造者模式是在当创建复杂对象的算法应该独立于该对象的组成部分以及它们的装配方式时使用的模式。