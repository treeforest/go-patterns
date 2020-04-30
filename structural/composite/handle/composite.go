package handle

import (
	"github.com/treeforest/go-patterns/structural/composite"
	"container/list"
	"fmt"
)

// 枝节点
type composite struct {
	name string
	children *list.List
}

func NewComposite(name string) component.Component {
	p := new(composite)
	p.name = name
	p.children = list.New()
	return p
}

func (p *composite) Add(c component.Component) error {
	p.children.PushBack(c)
	return nil
}

func (p *composite) Remove(c component.Component) error {
	for e := p.children.Front(); e != p.children.Back(); e = e.Next() {
		if e.Value == c {
			p.children.Remove(e)
			break
		}
	}
	return  nil
}

func (p *composite) Display(depth int) {
	var s string = ""
	for i := 0; i < depth; i++ {
		s += "-"
	}
	fmt.Println(s, p.name)

	for e := p.children.Front(); e != p.children.Back(); e = e.Next() {
		c := e.Value.(component.Component)
		c.Display(depth + 2)
	}
}

