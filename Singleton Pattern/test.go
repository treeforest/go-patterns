package main

import (
	"github.com/treeforest/design-mode/singleton/example1"
	"github.com/treeforest/design-mode/singleton/example2"
	"github.com/treeforest/design-mode/singleton/example3"
	"github.com/treeforest/design-mode/singleton/example4"
	"time"
	"fmt"
)

func main(){
	Test_example1()
	Test_example2()
	Test_example3()
	Test_example4()
}

func Test_example1(){
	fmt.Println("Test example1:")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p ", example1.GetInstance())
		}()
	}
	time.Sleep(time.Second*1)
}

func Test_example2(){
	fmt.Println("\nTest example2:")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p ", example2.GetInstance())
		}()
	}
	time.Sleep(time.Second*1)
}

func Test_example3(){
	fmt.Println("\nTest example3:")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p ", example3.GetInstance())
		}()
	}
	time.Sleep(time.Second*1)
}


func Test_example4(){
	fmt.Println("\nTest example4:")
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p ", example4.GetInstance())
		}()
	}
	time.Sleep(time.Second*1)
}