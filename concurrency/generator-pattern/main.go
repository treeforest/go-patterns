package main

import (
	"fmt"
	"time"
)

// 生成器实现：生成一个值序列
func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// 其他操作
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}

func main() {
	ch := Count(1, 20)

	for {
		select {
		case n, ok := <- ch:
			if !ok {
				break
			}
			fmt.Printf("%d ", n)
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
