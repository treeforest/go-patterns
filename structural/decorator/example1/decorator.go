package main

import "fmt"

type Object func(int) int

func CalculateDecorate(o Object) Object {
	return func(n int) int {
		fmt.Println("Starting the execution with the integer : ", n)

		result := o(n)

		fmt.Println("Execution is completed  with the result : ", result)

		return result
	}
}

func Double(n int) int {
	return n * 2
}

func Square(n int) int {
	return n * n
}

func main() {
	o := CalculateDecorate(Double)
	o(10)

	o = CalculateDecorate(Square)
	o(10)
}
