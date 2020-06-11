# Generator Pattern

生成器模式: 将一系列数据的生成放在不同的函数中实现。

# Implementation

```
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
```

# Usage

```
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
```

```
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 
```