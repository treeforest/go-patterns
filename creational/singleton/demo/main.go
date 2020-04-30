package main

import (
	"fmt"
	"github.com/treeforest/go-patterns/creational/singleton/singleton"
)

func main() {
	s := singleton.New()
	s["this"] = "singleton"

	s2 := singleton.New()

	fmt.Println("This is", s2["this"])
}
