package squarefactory

import (
	"github.com/treeforest/design-pattern/factory/factorymethod2/shape"
	"github.com/treeforest/design-pattern/factory/factorymethod2/shape/square"
)

type SquareFactory struct {
	// ...
}

func (p *SquareFactory) CreateSquare() shape.Shape{
	return new(square.Square)
}
