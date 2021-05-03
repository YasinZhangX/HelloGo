package main

import (
	"fmt"
	"time"
)

func main() {
	channelDemo()
	fmt.Println("------------------")
	bufferedChannelDemo()
	fmt.Println("------------------")
	channelClosed()
}

func channelClosed() {
	c := make(chan int, 3)

	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)

	time.Sleep(50 * time.Millisecond)
}

func bufferedChannelDemo() {
	c := make(chan int, 3)

	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'

	time.Sleep(50 * time.Millisecond)
}

func channelDemo() {
	var channels [10]chan int // c == nil
	// c := make(chan int)

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(50 * time.Millisecond)
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
