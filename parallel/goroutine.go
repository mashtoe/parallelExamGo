package main

import (
	"fmt"
	"time"
)

func main(){
	printWord("Sequential")

	go printWord("Goroutine")

	printWord("Hello")
}

func printWord(word string){
	for i:= 0; i < 3 ; i++{
		fmt.Println(word, ":", (i+1))
		time.Sleep(time.Second * 1)
	}
}