package product02

import (
		"fmt"
)

type circle struct {

}

func (p *circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}
