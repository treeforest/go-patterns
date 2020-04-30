package car

import (
	"fmt"
	"github.com/treeforest/go-patterns/creational/builder"
)

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
