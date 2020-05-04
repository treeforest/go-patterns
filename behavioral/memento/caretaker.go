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
