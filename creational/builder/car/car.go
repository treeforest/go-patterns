package car

import "fmt"

type car struct {
	Interface
	color  Color
	wheels Wheels
	speed  Speed
}

func (c *car) Drive() error {
	fmt.Printf("A %s car with %s tires is going at %f ", c.color, c.wheels, c.speed)
	if c.speed == MPH {
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
