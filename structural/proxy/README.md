# Proxy Pattern

代理模式：为其他对象提供一种代理以控制对这个对象的访问。

# Implementation

## Types

```
package proxy

// 若要使用代理功能，则必须实现该接口的相同方法
type IObject interface {
	ObjDo(action string)
}
```

## Different Implementations

```
// 实际进行数据处理的类
type object struct {
	//...
}

// ObjDo是实现于IObject的接口
// 实现了真正的逻辑处理
func (obj *object) ObjDo(action string) {
	fmt.Printf("I can, %s", action)
}
```

```
// 具备有拦截行为的代理类
type proxyObject struct {
	once sync.Once
	obj  *object
}

func NewProxyObject() proxy.IObject {
	return new(proxyObject)
}

// ObjDo是实现于IObject的接口
// 目的是为了拦截访问真正的Object的行为
func (p *proxyObject) ObjDo(action string) {
	p.once.Do(func() {
		p.obj = new(object)
	})

	// do something(拦截处理等)

	p.obj.ObjDo(action)
}
```

# Usage

```
proxy := handle.NewProxyObject()
proxy.ObjDo("Run")
```

```
I can, Run
```

# Rules of Thumb

第一：远程代理，也就是为一个对象在不同的地址空间提供局部代表，这样可以隐藏一个对象存在不同地址空间的事实。  
第二：虚拟代理：是根据需要创建开销很大的对象。通过它来存放实例化需要很长时间的真实对象。  
第三：安全代理：用来控制真实对象访问时的权限。  
第四：智能指引：是指当调用真实的对象时，代理处理另外一些事。