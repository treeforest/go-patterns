package main

type ComputerPartVisitor interface {
	Visit(ComputerPart)
}

type ComputerPartDisplayVisitor struct {

}

type (
	ComputerPart interface {
		Accept(vistor ComputerPartVisitor)
	}

	Computer struct {
		parts []ComputerPart
	}

	KeyBoard struct {}
	Monitor struct {}
	Mouse struct {}
)

func CreateComputer() ComputerPart {
	c := new(Computer)
	c.parts = make([]ComputerPart, 3)
	append(c.parts, new(Mouse))
	append(c.parts, new(Monitor))
	append(c.parts, new(KeyBoard))
	return c
}

func (c *Computer) Accept(vistor ComputerPartVisitor) {
	for i,_ := range c.parts{
		c.parts[i].Accept(vistor)
	}

}

func (m *Mouse) Accept(vistor ComputerPartVisitor) {

}

func (m *Monitor) Accept(vistor ComputerPartVisitor) {

}

func (k *KeyBoard) Accept(vistor ComputerPartVisitor) {

}

func main() {

}
