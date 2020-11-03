package main

func main() {
	var abstraction Abstraction

	abstraction = &RefinedAbstraction{
		imp: new(ConcreteImplementorA),
	}
	abstraction.Operation()

	abstraction = &RefinedAbstraction{
		imp: new(ConcreteImplementorB),
	}
	abstraction.Operation()
}
