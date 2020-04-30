package proxy

// 若要使用代理功能，则必须实现该接口的相同方法
type IObject interface {
	ObjDo(action string)
}
