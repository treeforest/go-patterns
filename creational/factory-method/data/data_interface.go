package data

import "io"

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}
