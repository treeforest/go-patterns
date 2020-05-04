package handle

import "github.com/treeforest/go-patterns/behavioral/chain-of-responsibility"

// 请求
type request struct {
	typ     string
	num     int
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
