package handler

type Manager interface {
	SetSuperior(superior Manager)
	RequestApplications(request Request)
}

type Request interface {
	GetRequestType() string
	GetRequestContent() string
	GetRequestNumber() int
}
