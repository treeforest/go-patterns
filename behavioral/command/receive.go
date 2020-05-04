package command

import "fmt"

type Receiver interface {
	Action() error
}

type concreteReceiver struct {
}

func CreateReceiver() Receiver {
	return new(concreteReceiver)
}

func (r *concreteReceiver) Action() error {
	// 是否否决该命令
	fmt.Println("执行请求！")
	return nil
}
