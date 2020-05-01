package product03

import "fmt"

type square struct {
}

func (p *square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}
