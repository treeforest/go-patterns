package main

import (
	"path/filepath"
	"os"
	"errors"
	"crypto/md5"
	"io/ioutil"
)

type result struct {
	path string
	sum [md5.Size]byte
	err error
}

func walkFiles(done <-chan struct{}, root string) (<-chan string, <- chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		// 在遍历结束之后，关闭paths管道
		defer close(paths)

		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
				case paths <- path:
				case <- done:
					return errors.New("walk canceled.")
			}
			return nil
		})
	}()

	return paths, errc
}


func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := ioutil.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}
		case <- done:
			return
		}
	}
}