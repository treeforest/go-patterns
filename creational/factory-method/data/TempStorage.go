package data

import (
	"io"
	"github.com/treeforest/go-patterns/creational/factory-method"
)

type tempStorage struct {
}

func newTempStorage() (fm.Store, error) {
	return new(tempStorage), nil
}

func (t *tempStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return rwc, err
}
