package data

import (
	"github.com/treeforest/go-patterns/creational/factory-method"
	"io"
)

type diskStorage struct {
}

func newDiskStorage() (fm.Store, error) {
	return new(diskStorage), nil
}

func (p *diskStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return rwc, err
}
