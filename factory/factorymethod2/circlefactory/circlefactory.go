package circlefactory

import (
	"github.com/treeforest/design-pattern/factory/factorymethod2/shape"
	"github.com/treeforest/design-pattern/factory/factorymethod2/shape/circle"
)

type CircleFactory struct {
	//...
}

func (p *CircleFactory) CreateCircle() shape.Shape{
	return new(circle.Circle)
}
