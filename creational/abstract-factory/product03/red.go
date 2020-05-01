package product03

import "fmt"

type red struct {
}

func (p *red) Fill() {
	fmt.Println("Inside Red::fill() method.")
}
