package main

import (
	"github.com/treeforest/go-patterns/creational/factory-method/data"
	"fmt"
)

func main() {
	store, _ := data.NewStore(data.MemoryStorage)
	f, _ := store.Open("file")
	defer f.Close()

	f.Write([]byte("Hello,"))
	f.Write([]byte(" factory method."))

	s := make([]byte, 100)
	n, _ := f.Read(s)

	fmt.Println(string(s[0:n]))
}
