package main

import (
	"fmt"
	"sync"
	"time"
)
const concurrent_summers = 4

func sum_channel(ch chan int) int {
	total := 0
	for val := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("got value from input",val)
		total += val
	}
	fmt.Println(total)
	return total
}

func main() {
	nums := []int{1, 2, 3,4,100,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200,200}
	input := make(chan int)
	output := make(chan int, concurrent_summers)

	// feed starting values into the input channel
	go func() {
		for i:= 0;i< len(nums) ;i++ {
			input <- nums[i]
			// time.Sleep(time.Second * 2)
			fmt.Println("input: ", nums[i])
		}
		println("Closing")
		close(input)
	}()
	var wg sync.WaitGroup
	wg.Add(concurrent_summers)

	// start summers which will work independently and send their results
	// to the output channel
	for i := 0; i < concurrent_summers; i++ {
		go func() {
			output <- sum_channel(input)
			wg.Done()
		}()
	}
	wg.Wait()
	close(output)
	// now sum the 4 results in the output channel
	//fmt.Println("Total is:", sum_channel(output))
}