package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	sw      sync.WaitGroup
	initNum int64
)

func accumulation() {
	defer sw.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&initNum, 1)
	}
}

func main() {

	for i := 0; i < 10; i++ {
		sw.Add(1)
		go accumulation()
	}
	sw.Wait()
	fmt.Println("initNum最终结果：", initNum)
}
