package data

import "io"

type tempStorage struct {
	Store
}

func newTempStorage() (Store, error) {
	return new(tempStorage), nil
}

func (t *tempStorage) Open(f string) (rwc io.ReadWriteCloser, err error) {
	return rwc, err
}
