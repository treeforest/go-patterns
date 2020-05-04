package main

import "github.com/treeforest/go-patterns/behavioral/command"

func main() {
	r := command.CreateReceiver()
	c := command.CreateCommand(r)
	i := command.CreateInvoker()

	i.SetCommand(c)
	i.ExecuteCommand()
}
