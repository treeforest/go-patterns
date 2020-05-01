package product01

import (
	"fmt"
)

type rectangle struct {
}

func (p *rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}
