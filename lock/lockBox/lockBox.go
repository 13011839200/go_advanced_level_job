package main

import (
	"fmt"
	"sync"
)

var (
	sm      sync.Mutex
	sw      sync.WaitGroup
	initNum int
)

func accumulation() {
	defer sw.Done()
	for i := 0; i < 1000; i++ {
		sm.Lock()
		initNum++
		sm.Unlock()
	}
}

func main() {
	for i := 0; i < 10; i++ {
		sw.Add(1)
		go accumulation()
	}
	sw.Wait()
	fmt.Println("initNum的最终结果为：", initNum)

}
