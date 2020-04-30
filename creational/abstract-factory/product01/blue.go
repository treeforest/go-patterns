package product01

import (
	"fmt"
	)

type blue struct {

}

func (p *blue) Fill() {
	fmt.Println("Inside Blue::fill() method.")
}
