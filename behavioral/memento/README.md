# Memento Pattern

备忘录模式：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。

# Implementation

```$xslt
package memento

import "fmt"

type Originator interface {
	SetState(state string)
	GetState() string
	CreateMemento() Memento
	SetMemento(m Memento)
	Show()
}

/*
 * 发起人
 *
 * 负责创建一个备忘录 Memento，用以记录当前时刻它的内部状态，并可使用备忘录恢复到内部状态。
 */
type originator struct {
	state string // 状态属性
}

func CreateOriginator() Originator {
	return new(originator)
}

func (o *originator) SetState(state string) {
	o.state = state
}

func (o *originator) GetState() string {
	return o.state
}

// 创建备忘录，将当前需要保存的信息导入并实例化出一个Memento对象
func (o *originator) CreateMemento() Memento {
	return createMemento(o.state)
}

// 恢复备忘录，将Memento导入并将相关数据恢复
func (o *originator) SetMemento(m Memento) {
	o.state = m.GetState()
}

// 显示数据
func (o *originator) Show() {
	fmt.Println("state = ", o.state)
}
```

```$xslt
package memento

type Memento interface {
	GetState() string
}

/*
 * 备忘录
 *
 * 负责存储 Originator 对象的内部状态，并可防止 Originator 以外的其他对象访问备忘录 Memento
 */
type memento struct {
	state string
}

// 不导出
func createMemento(state string) Memento {
	m := new(memento)
	m.state = state
	return m
}

func (m *memento) GetState() string {
	return m.state
}
```

```$xslt
package memento

type Caretaker interface {
	GetMemento() Memento
	SetMemento(m Memento)
}

/*
 * 管理者类
 *
 * 负责保存好备忘录 Memento
 */
type caretaker struct {
	m Memento
}

func CreateCaretaker() Caretaker {
	return new(caretaker)
}

func (c *caretaker) GetMemento() Memento {
	return c.m
}

func (c *caretaker) SetMemento(m Memento) {
	c.m = m
}
```

# Usage

```$xslt
o := memento.CreateOriginator()
o.SetState("On") // 设置Originator初始状态属性为“On”
o.Show()

c := memento.CreateCaretaker()
c.SetMemento(o.CreateMemento()) // 保存状态时，由于有了很好的封装，可以隐藏Originator的实现细节

o.SetState("Off") // Originator改变了状态属性为“Off”
o.Show()

o.SetMemento(c.GetMemento()) // 恢复上一状态
o.Show()
```

```$xslt
state =  On
state =  Off
state =  On
```

# Rules of Thumb 

Memento 模式比较适用于功能比较复杂的，但需要维护或记录属性历史的类，  
或者需要保存的属性只是众多属性中的一小部分时，Originator 可以根据保存的 Memento 信息还原到前一状态。