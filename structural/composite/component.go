package component

// 组合中所有对象的抽象接口（透明方式）
type Component interface {
	Add(c Component) error
	Remove(c Component) error
	Display(depth int)
}
