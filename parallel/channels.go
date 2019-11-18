package main

import "fmt"

func main(){
	ch := make(chan int)
	resch := make(chan int)

	go AddOne(ch, 9)
	go MulBy10(ch, resch)

	result := <-resch
	fmt.Println("Result:", result)
}

func AddOne(ch chan<- int, i int) {
	i++
	ch <- i
}

func MulBy10(ch <-chan int, resch chan<- int) {
	i := <-ch
	i *= 10
	resch <- i
}