package main

import (
	"fmt"
)

func main() {
	runChannelWithSingleValue()
}

func runChannelWithSingleValue() {
	fmt.Println("Run channel with single value")
	var c = make(chan int) // create channel that can holds an int value
	// deadLockExample(c)
	go process(c)
	// note that there is no WaitGroup. so code continues and doesn't wait for the process routine to finish
	fmt.Println(<-c) // this code is blocked (waiting for some value be added to channel c). process function will concurrently unblock this code when c <- i happens
	for i := range c {
		fmt.Println(i) // this code is blocked (waiting for some value be added to channel c)
	}
}

func process(c chan int) {
	// defer means do this right before this function exits
	defer close(c) // notifies any other process using this channel we are done and the main function will break out of the loop (for i := range c) and exit
	for i := 0; i < 5; i++ {
		c <- i
	}
}

func deadLockExample(c chan int) {
	c <- 1 // c : [1]. adds value 1 to the channel
	// code is blocked until some one reads from channel
	// Since we only read value from channel in next line, the code will never be unblocked. Then, fatal error: all goroutines are asleep - deadlock!
	var i = <-c // c : [] and i = 1. remove value 1 from channel
	fmt.Println(i)
}
