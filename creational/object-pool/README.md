# Object Pool Pattern

对象池模式

# Implementation

```
package pool

import "fmt"

type Object struct {
	id int
}

func (o *Object) Do(s string) {
	fmt.Printf("ID = %d, Object Do %s\n", o.id, s)
}
```

```
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
```

# Usage

```
p := pool.New(4)

for {
    select {
    case obj := <-*p:
        time.Sleep(time.Second * 1)
        obj.Do("Hello")
        *p <- obj
    default:
        return
    }
}
```

```
ID = 0, Object Do Hello
ID = 1, Object Do Hello
ID = 2, Object Do Hello
ID = 3, Object Do Hello
ID = 0, Object Do Hello
ID = 1, Object Do Hello
ID = 2, Object Do Hello
...
```

# Rules of Thumb

1、适用于当初始化一个对象实例比维护一个对象实例的开销更大的情形  
2、由于对象是预先初始化的，它对性能有积极的影响  
3、如果不是稳定的需求且需求会出现峰值，则需要进行维护的开销可能会超过了对象池的好处。