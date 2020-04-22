package shapefactory

import (
	"github.com/treeforest/design-pattern/factory/factorymethod1/shape/rectangle"
	"github.com/treeforest/design-pattern/factory/factorymethod1/shape/circle"
	"github.com/treeforest/design-pattern/factory/factorymethod1/shape/square"
	"github.com/treeforest/design-pattern/factory/factorymethod1/shape"
	"sync"
)

// 图形工厂
type shapeFactory struct {
	//..
}

var once sync.Once
var instance *shapeFactory = nil
func CreateShapeFactory() *shapeFactory{
	once.Do(func() {
		instance = new(shapeFactory)
	})
	return instance
}

func (p *shapeFactory)CreateRectangle() shape.Shape{
	//...
	return new(rectangle.Rectangle)
}

func (p *shapeFactory) CreateCircle() shape.Shape {
	//...
	return new(circle.Circle)
}

func (p *shapeFactory) CreateSquare() shape.Shape {
	//...
	return new(square.Square)
}
