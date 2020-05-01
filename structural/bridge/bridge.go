package bridge

type Soft interface {
	Run() error
}

type Brand interface {
	SetSoft(soft Soft)
	Run()
}
