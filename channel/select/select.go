package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	var c1, c2 = generator(), generator()
	worker := createWorker(0)

	var values []int
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len: ", len(values))
		case <-tm:
			fmt.Println("end")
			return
		}
	}
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
