package car

type carBuilder struct {
	Builder
	c *car
}

func NewCarBuilder() Builder {
	cb := new(carBuilder)
	cb.c = new(car)
	return cb
}

func (cb *carBuilder) Color(c Color) Builder {
	cb.c.color = c
	return cb
}

func (cb *carBuilder) Wheels(w Wheels) Builder {
	cb.c.wheels = w
	return cb
}

func (cb *carBuilder) TopSpeed(s Speed) Builder {
	cb.c.speed = s
	return cb
}

func (cb *carBuilder) Build() Interface {
	return cb.c
}
