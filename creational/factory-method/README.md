# Factory Method Pattern

工厂方法模式：定义一个用于创建对象的接口，让子类决定实例化哪个类，工厂方法使一个类的实例化延迟到其子类。

# Implementation

示例实现了对不同缓存对象的使用

## Types

```
package data

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}
```

## Different Implementations

```
type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) (s fm.Store, err error) {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return newDiskStorage()
	default:
		return newTempStorage()
	}
}
```

```
type memory struct {
	io.ReadWriteCloser
	buf []byte
}

func (m *memory) Read(p []byte) (n int, err error) {
	for i, v := range m.buf {
		p[i] = v
	}
	return len(p), nil
}

func (m *memory) Write(p []byte) (n int, err error) {
	m.buf = append(m.buf, p...)
	return len(p), nil
}

func (m *memory) Close() error {
	m.buf = m.buf[0:0]
	return nil
}

type memoryStorage struct {
}

func newMemoryStorage() (fm.Store, error) {
	return new(memoryStorage), nil
}

func (m *memoryStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return new(memory), nil
}
```

```
type diskStorage struct {
}

func newDiskStorage() (fm.Store, error) {
	return new(diskStorage), nil
}

func (p *diskStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return rwc, err
}
```

```
type tempStorage struct {
}

func newTempStorage() (fm.Store, error) {
	return new(tempStorage), nil
}

func (t *tempStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return rwc, err
}
```

# Usage

```
store, _ := data.NewStore(data.MemoryStorage)
f, _ := store.Open("file")
defer f.Close()

f.Write([]byte("Hello,"))
f.Write([]byte(" factory method."))

s := make([]byte, 100)
n, _ := f.Read(s)

fmt.Println(string(s[0:n]))
```

```
Hello, factory method.
```

# Rules of Thumb

有利于产品族内产品的增加，不利于产品族的增加。