package main

import "github.com/treeforest/go-patterns/behavioral/memento"

func main() {
	o := memento.CreateOriginator()
	o.SetState("On") // 设置Originator初始状态属性为“On”
	o.Show()

	c := memento.CreateCaretaker()
	c.SetMemento(o.CreateMemento()) // 保存状态时，由于有了很好的封装，可以隐藏Originator的实现细节

	o.SetState("Off") // Originator改变了状态属性为“Off”
	o.Show()

	o.SetMemento(c.GetMemento()) // 恢复上一状态
	o.Show()
}
