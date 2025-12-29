package main

import "fmt"

func addten(integer *int) int {
	value := *integer + 10
	return value
}

func main() {
	val := 10
	result := addten(&val)
	fmt.Println(result)
}
