package main

import "github.com/treeforest/go-patterns/structural/proxy/handle"

func main() {
	proxy := handle.NewProxyObject()
	proxy.ObjDo("Run")
}
