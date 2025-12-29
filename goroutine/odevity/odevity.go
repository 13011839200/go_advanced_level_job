package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

const (
	even_num string = "偶数"
	odd_num  string = "奇数"
)

func printOdevity(isEven bool) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		remainder := i % 2
		if remainder == 0 && isEven {
			fmt.Printf("%s:%d\n", even_num, i)
		}
		if remainder != 0 && !isEven {
			fmt.Printf("%s:%d\n", odd_num, i)

		}

	}
}

func main() {
	count := 2

	for i := 1; i <= count; i++ {
		wg.Add(1)
		// 第一个函数打印奇数，第二个函数打印偶数
		go printOdevity(i == 2)
	}
	wg.Wait()
}
