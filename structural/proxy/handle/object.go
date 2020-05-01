package handle

import (
	"fmt"
)

// 实际进行数据处理的类
type object struct {
	//...
}

// ObjDo是实现于IObject的接口
// 实现了真正的逻辑处理
func (obj *object) ObjDo(action string) {
	fmt.Printf("I can, %s", action)
}
