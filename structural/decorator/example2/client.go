package main

// Decorator 将请求转发给它的 Component 对象，并有可能在转发请求前后执行一些附加的动作
// 具体来说，就是我现在要给我的组件进行装饰（增加职责），但是要避免装饰物过多带来的类爆炸式增长。
func main() {
	c := new(ConcreteComponent)
	A := new(ConcreteDecoratorA)
	B := new(ConcreteDecoratorB)

	A.Dec.SetComponent(c)
	B.Dec.SetComponent(A)
	B.Operation()
}
