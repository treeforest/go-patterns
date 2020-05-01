# Chain of Responsibility Pattern

职责链模式：使多个对象都有机会处理请求，从而避免请求的发送者和接受者之间的耦合关系。将这个对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它2为止。

# Implementation

该示例实现了公司职员对加薪或请假请求的链式处理过程。

## Types

```
package handler

type Manager interface {
	SetSuperior(superior Manager)
	RequestApplications(request Request)
}

type Request interface {
	GetRequestType() string
	GetRequestContent() string
	GetRequestNumber() int
}
```

## Different Implementations

```
// 请求
type request struct {
	typ string
	num int
	content string
}

func CreateRequest(typ, content string, num int) handler.Request {
	r := new(request)
	r.typ = typ
	r.content = content
	r.num = num
	return r
}

func (r *request) GetRequestType() string {
	return r.typ
}

func (r *request) GetRequestContent() string {
	return r.content
}

func (r *request) GetRequestNumber() int {
	return r.num
}
```

```
// 管理者超类
type manager struct {
	handler.Manager
	superior handler.Manager
	name string
}

func (m *manager) SetSuperior(superior handler.Manager) {
	m.superior = superior
}
```

```
// 经理
type commonManager struct {
	manager
}

func CreateCommonManager(name string) handler.Manager {
	m := new(commonManager)
	m.name = name
	return m
}

func (m *commonManager) RequestApplications(request handler.Request) {
	if(request.GetRequestType() == "请假" && request.GetRequestNumber() <= 2) {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	}else {
		if m.superior != nil {
			// 将申请转交到上级
			m.superior.RequestApplications(request)
		}
	}
}
```

```
// 总监
type majordomo struct {
	manager
}

func CreateMajordomo(name string) handler.Manager {
	m := new(majordomo)
	m.name = name
	return m
}

func (m *majordomo) RequestApplications(request handler.Request) {
	if(request.GetRequestType() == "请假" && request.GetRequestNumber() <= 5) {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	}else {
		if m.superior != nil {
			// 将申请转交到上级
			m.superior.RequestApplications(request)
		}
	}
}
```

```
type generalManager struct {
	manager
}

func CreateGeneralManager(name string) handler.Manager {
	m := new(generalManager)
	m.name = name
	return m
}

func (m *generalManager) RequestApplications(request handler.Request) {
	if request.GetRequestType() == "请假" {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	} else if request.GetRequestType() == "加薪" && request.GetRequestNumber() <= 500 {
		fmt.Printf("%s:%s 数量%d 被批准。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	} else if request.GetRequestType() == "加薪" && request.GetRequestNumber() > 500{
		fmt.Printf("%s:%s 数量%d 再说吧。\n", m.name, request.GetRequestContent(), request.GetRequestNumber())
	}
}
```

# Usage

```
jinli := handle.CreateCommonManager("张三（经理）")
zongjian := handle.CreateMajordomo("李四（总监）")
zhongjingli := handle.CreateGeneralManager("王五（总经理）")

jinli.SetSuperior(zongjian)
zongjian.SetSuperior(zhongjingli)

// 客户端的申请都是由“经理”发起，实际谁来决策由具体管理类来处理，客户端不知道
req := handle.CreateRequest("请假","小菜请假", 1)
jinli.RequestApplications(req)

req2 := handle.CreateRequest("请假", "小菜请假", 4)
jinli.RequestApplications(req2)

req3 := handle.CreateRequest("加薪", "小菜请求加薪", 500)
jinli.RequestApplications(req3)

req4 := handle.CreateRequest("加薪", "小菜请求加薪", 1000)
jinli.RequestApplications(req4)
```

# Rules of Thumb

当心，一个请求极有可能到了链的末端都得不到处理，或者因为没有正确配置而得不到处理。
