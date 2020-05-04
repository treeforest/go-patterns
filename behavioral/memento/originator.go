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
