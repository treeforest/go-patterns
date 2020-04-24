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