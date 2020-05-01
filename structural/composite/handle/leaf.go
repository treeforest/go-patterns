package handle

import (
	"fmt"
	"github.com/treeforest/go-patterns/structural/composite"
)

// 叶子节点
type leaf struct {
	name string
}

func NewLeaf(name string) component.Component {
	p := new(leaf)
	p.name = name
	return p
}

func (p *leaf) Add(c component.Component) error {
	fmt.Println("Cannot add to a leaf.")
	return nil
}

func (p *leaf) Remove(c component.Component) error {
	fmt.Println("Cannot remove from a leaf.")
	return nil
}

func (p *leaf) Display(depth int) {
	var s string = ""
	for i := 0; i < depth; i++ {
		s += "-"
	}
	fmt.Println(s, p.name)
}
