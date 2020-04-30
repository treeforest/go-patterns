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
