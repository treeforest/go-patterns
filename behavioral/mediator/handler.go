package mediator

// 抽象中介者接口
type Mediator interface {
	Send(msg string, c Colleague) error
}

// 抽象同事接口
type Colleague interface {
	Send(msg string) error
	Notify(msg string) error
}
