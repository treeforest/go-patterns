package main

func Double(n int) int {
	return n * 2
}

func Square(n int) int {
	return n * n
}

func main () {
	o := CalculateDecorate(Double)
	o(10)

	o = CalculateDecorate(Square)
	o(10)
}

