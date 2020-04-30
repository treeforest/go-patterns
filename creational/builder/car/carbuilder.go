package car

import "github.com/treeforest/go-patterns/creational/builder"

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
