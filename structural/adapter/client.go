package main

func main() {
	var target Target
	target = &Adapter{
		Adp: new(Adaptee),
	}

	target.Request()
}
