package pool

import "fmt"

type Object struct {
	id int
}

func (o *Object) Do(s string) {
	fmt.Printf("ID = %d, Object Do %s\n", o.id, s)
}