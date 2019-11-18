package main

import (
	"fmt"
	"time"
)

func main() {
	c:= make(chan int)

	go add(2, 2, c)

	result := <-c

	fmt.Println(result)
}

func add(x int, y int, c chan int) {
	time.Sleep(time.Second)
	result := x + y
	c <- result
}