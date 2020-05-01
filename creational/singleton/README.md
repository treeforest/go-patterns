# Singleton Pattern

单例模式：保证一个类仅有一个实例，并提供一个访问它的全局访问点。

# Implementation

```
package singleton

import "sync"

type singleton map[string]string

var (
	once sync.Once

	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})

	return instance
}
```

# Usage

```
s := singleton.New()
s["this"] = "singleton"

s2 := singleton.New()

fmt.Println("This is", s2["this"])
```