package pool

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		o := new(Object)
		o.id = i
		p <- o
	}

	return &p
}
