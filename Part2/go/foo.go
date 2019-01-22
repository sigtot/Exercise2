package main

import (
    "fmt"
    "runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func numberServer(addNumber <-chan int, control <-chan int, number chan<- int) {
    i := 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		case n := <-addNumber:
		    i += n
        case c := <-control:
            switch c {
            case GetNumber:
                number<- i
            case Exit:
                return
            }
		}
	}
}

func incrementing(addNumber chan<-int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		addNumber <- 1
	}
	finished <- true
}

func decrementing(addNumber chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		addNumber <- -1
	}
    finished<- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
    addNumber := make(chan int)
    control := make(chan int)
    number := make(chan int)
    finished := make(chan bool)

    go numberServer(addNumber, control, number)

    go incrementing(addNumber, finished)
    go decrementing(addNumber, finished)
    <-finished
    <-finished

	control<-GetNumber
	fmt.Println("The magic number is:", <- number)
	control<-Exit
}
