# Command Pattern

命令模式：将一个请求封装为一个对象，从而使你可用不同的请求对客户参数化；对排队请求或记录请求日志，以及支持可撤销的操作。

# Implementation

```
package command

import "fmt"

type Receiver interface {
	Action() error
}

type concreteReceiver struct {

}

func CreateReceiver() Receiver{
	return new(concreteReceiver)
}

func (r *concreteReceiver) Action() error {
	// 是否否决该命令
	fmt.Println("执行请求！")
	return nil
}
```

```
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
```

```
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
```

# Usage

```$xslt
r := command.CreateReceiver()
c := command.CreateCommand(r)
i := command.CreateInvoker()

i.SetCommand(c)
i.ExecuteCommand()
```

```$xslt
执行请求！
```

# Rules of Thumb

第一：它能较容易地设计一个命令队列；  
第二：在需要的情况下，可以较容易地将命令记入日志；  
第三：允许接受请求的一方决定是否要否决请求；  
第四：可以容易地实现对请求的撤销和重做；  
第五：由于加进新的具体命令类不影响其他的类，因此增加新的具体命令类很容易。  
  
命令模式把请求一个操作对象与知道怎么执行一个操作的对象分割开。  

