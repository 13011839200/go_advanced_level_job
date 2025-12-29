package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sw sync.WaitGroup
)

func accumulation(addNum []int) {
	defer sw.Done()
	var total int
	start := time.Now
	defer func() {
		tim := time.Since(start())
		fmt.Printf("任务：%v的累加结果为：%v，耗时为：%vμs\n", addNum, total, tim.Microseconds())
	}()
	if addNum == nil {
		return
	}
	for _, num := range addNum {
		total += num
	}

}

func main() {

	slices1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slices2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slices3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	sw.Add(3)
	go accumulation(slices1)
	go accumulation(slices2)
	go accumulation(slices3)
	sw.Wait()

}
