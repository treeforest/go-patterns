package main

import "github.com/treeforest/go-patterns/structural/composite/handle"

func main() {
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
}
