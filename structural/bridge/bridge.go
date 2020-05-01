package bridge

type ISoft interface {
	Run() error
}

type IBrand interface {
	SetSoft(soft ISoft)
	Run()
}
