package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func main() {
	channelDemo()
}

func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() { wg.Done() },
	}

	go doWork(id, w)

	return w
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
	}
}
