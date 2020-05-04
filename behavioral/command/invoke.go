package command

type Invoker interface {
	SetCommand(cmd Commander)
	ExecuteCommand()
}

// 实现对命令的接收和执行
// 使用链表保存命令，实现连续的执行一系列命令（命令队列）
type invoke struct {
	cmd Commander
	// cmds *list.List
}

func CreateInvoker() Invoker {
	return new(invoke)
}

func (p *invoke) SetCommand(cmd Commander) {
	p.cmd = cmd
}

func (p *invoke) ExecuteCommand() {
	// 记入日志
	p.cmd.Execute()
}