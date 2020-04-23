package main

import (
	"github.com/treeforest/go-patterns/creational/object-pool/pool"
	"time"
)

func main() {
	p := pool.New(4)

	for{
		select {
		case obj := <- *p :
			time.Sleep(time.Second * 1)
			obj.Do("Hello")
			*p <- obj
		default:
			return
		}
	}
}

