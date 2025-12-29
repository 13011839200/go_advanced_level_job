package main

import (
	"fmt"
	"sync"
)

type util struct {
	nums []int
	cha  chan int
	sw   sync.WaitGroup
}

func sendNum(obj *util) {
	defer obj.sw.Done()
	for _, num := range obj.nums {
		obj.cha <- num
	}
	close(obj.cha)
}

func receiveNum(obj *util) {
	defer obj.sw.Done()

	for num := range obj.cha {
		fmt.Println(num)

	}

	// for {
	// 	select {
	// 	case num, ok := <-obj.cha:
	// 		if !ok {
	// 			return
	// 		}
	// 		fmt.Println(num)

	// 	default:
	// 		fmt.Println("正在接受中。。。")
	// 	}

	// }
}

func main() {
	obj := util{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, make(chan int), sync.WaitGroup{}}

	obj.sw.Add(2)

	go sendNum(&obj)
	go receiveNum(&obj)

	obj.sw.Wait()

}
