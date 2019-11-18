package main

import "fmt"

func main(){
	ch := make(chan int)
	ch2 := make(chan int)

	go addOne(ch, 9)
	go mulBy10(ch, ch2)

	result := <-ch2
	fmt.Println("Result:", result)
}

func addOne(ch chan<- int, i int) {
	i++
	ch <- i
}

func mulBy10(ch <-chan int, resch chan<- int) {
	i := <-ch
	i *= 10
	resch <- i
}