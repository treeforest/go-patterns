package main

import (
	"fmt"
	"github.com/treeforest/go-patterns/creational/abstract-factory/product01"
	"github.com/treeforest/go-patterns/creational/abstract-factory/product02"
	"github.com/treeforest/go-patterns/creational/abstract-factory/product03"
)

func main() {
	fmt.Println("product01: ")
	prd01 := product01.CreateProduct01Factory()
	c01 := prd01.CreateColor()
	c01.Fill()
	s01 := prd01.CreateShape()
	s01.Draw()

	fmt.Println("\nproduct02: ")
	prd02 := product02.CreateProduct02Factory()
	c02 := prd02.CreateColor()
	c02.Fill()
	s02 := prd02.CreateShape()
	s02.Draw()

	fmt.Println("\nproduct03: ")
	prd03 := product03.CreateProduct03Factory()
	c03 := prd03.CreateColor()
	c03.Fill()
	s03 := prd03.CreateShape()
	s03.Draw()
}
