# Composite Pattern

组合模式：将对象组合成树形结构以表示‘部分-整体’的层次结构，组合模式使得用户对单个对象和组合对象的使用具有一致性。

# Implementation

示例展示了组合模式（透明方式）的实现。

## Types

```
package component

// 组合中所有对象的抽象接口（透明方式）
type Component interface {
	Add(c Component) error
	Remove(c Component) error
	Display(depth int)
}
```

## Different Implementations

```
// 枝节点
type composite struct {
	name     string
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
	return nil
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
```

```
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
```

# Usage

```
root := handle.NewComposite("root")
root.Add(handle.NewLeaf("Leaf A"))
root.Add(handle.NewLeaf("Leaf B"))

comp := handle.NewComposite("Composite X")
comp.Add(handle.NewLeaf("Leaf XA"))
comp.Add(handle.NewLeaf("Leaf XB"))

root.Add(comp)

comp2 := handle.NewComposite("Composite XY")
comp2.Add(handle.NewLeaf("Leaf XYA"))
comp2.Add(handle.NewLeaf("Leaf XYB"))

comp.Add(comp2)

root.Add(handle.NewLeaf("Leaf C"))

leaf := handle.NewLeaf("Leaf D")
root.Add(leaf)
root.Remove(leaf)

root.Display(1)
```

```
- root
--- Leaf A
--- Leaf B
--- Composite X
----- Leaf XA
----- Leaf XB
--- Leaf C
```

# Rules of Thumb

当你发现需求中是体现部分与整体层次的结构时，以及你希望用户可以忽略组合对象与单个对象的不同，统一的使用组合结构中的对象时，就应该考虑用组合模式了。