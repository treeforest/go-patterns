package command

type Commander interface {
	Execute() error
}

// 基类，一个命令绑定一个命令接收者（执行对象）
type command struct {
	receiver Receiver
}

func (c *command) Execute() error {
	return c.receiver.Action()
}

// 实体类
type concreteCommand struct {
	command
	//...
}

func CreateCommand(receiver Receiver) Commander {
	c := new(concreteCommand)
	c.receiver = receiver // 绑定命令的执行对象
	return c
}