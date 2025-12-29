package main

import (
	"fmt"
	"sync"
)

var sw sync.WaitGroup

func sendNum(cha chan<- int) {
	defer sw.Done()
	for i := 0; i < 100; i++ {
		cha <- i
		fmt.Println("<----:", i)
	}
	close(cha)
}

func receiveNum(cha <-chan int) {
	defer sw.Done()

	for num := range cha {

		fmt.Println("---->:", num)

	}

}

func main() {
	cha := make(chan int, 10)
	sw.Add(2)
	go sendNum(cha)
	go receiveNum(cha)
	sw.Wait()
}
