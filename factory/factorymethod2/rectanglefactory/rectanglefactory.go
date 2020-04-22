package rectanglefactory

import (
	"github.com/treeforest/design-pattern/factory/factorymethod2/shape"
	"github.com/treeforest/design-pattern/factory/factorymethod2/shape/rectangle"
)

type RectangleFactory struct {
	// ...
}

func (p *RectangleFactory) CreateRectangle() shape.Shape{
	return new(rectangle.Rectangle)
}