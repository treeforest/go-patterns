package handle

import (
	"github.com/treeforest/go-patterns/structural/proxy"
	"sync"
)

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
