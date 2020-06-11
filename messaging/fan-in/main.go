package main

import (
	"sync"
	"fmt"
	"time"
)

// 合并不同的 channels 到一个 channel 中
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// 为每一个 cs 中的管道开启一个 send 协程，
	// send 从 c 中拷贝值到 out 中，直到 c 关闭后，再调用 wg.Done.
	send := func(c <- chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// 直到左右的 send 协程执行完毕后，再进行关闭。
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	defer func() {
		close(c1)
		close(c2)
		close(c3)
	}()

	f := func(start, end int, c chan int) {
		for i := start; i <= end; i++ {
			c <- i
		}
	}

	go f(1, 5, c1)
	go f(6,10, c2)
	go f(11, 15, c3)

	ch := Merge(c1, c2, c3)
	for {
		select {
			case n := <-ch:
				fmt.Printf("%d ", n)
			default:
			 	time.Sleep(time.Millisecond * 10)
		}
	}
}