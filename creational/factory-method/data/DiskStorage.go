package data

import (
	"io"
)

type diskStorage struct {
	Store
}

func newDiskStorage() (Store, error) {
	return new(diskStorage), nil
}

func (p *diskStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return rwc, err
}
