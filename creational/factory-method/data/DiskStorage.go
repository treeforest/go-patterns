package data

import (
	"io"
	"github.com/treeforest/go-patterns/creational/factory-method"
)

type diskStorage struct {
}

func newDiskStorage() (fm.Store, error) {
	return new(diskStorage), nil
}

func (p *diskStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return rwc, err
}
