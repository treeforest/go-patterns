package product02

import (
	"fmt"
	)

type green struct {

}

func (p *green) Fill() {
	fmt.Println("Inside Green::fill() method.")
}
