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
