package main

import "time"

var sem = make(chan int, 4)
func main() {
	for i:= 0; i < 20; i++ {
		sem <- 1
		temp:=i
		go func(){
			SlowOperation((temp+1));
			<-sem
		}()
	}
}

func SlowOperation(num int){
	println("running: ", num)
	time.Sleep(time.Second * 10);
}