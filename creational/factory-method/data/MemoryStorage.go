package data

import (
	"io"
)

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
	Store
}

func newMemoryStorage() (Store, error) {
	return new(memoryStorage), nil
}

func (m *memoryStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return new(memory), nil
}
