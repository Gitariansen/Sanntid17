package main

import (
	"fmt"
	"runtime"
)

var i int = 0

func thread_1_routine(ch, finish chan int) {
	
	for x := 0; x <= 1000001; x++ {
		<- ch
		i++
		ch <- 1
		
	}
	finish <- 1
}

func thread_2_routine(ch, finish chan int) {
	
	for x := 0; x <= 1000000; x++ {
		<- ch		
		i--
		ch <- 1
	}
	finish <- 1	
}


func main() { 
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int, 1)
	ch <- 1
	finish := make(chan int, 1)
	go thread_1_routine(ch, finish)
	go thread_2_routine(ch, finish)
	<- finish
	<- finish

	fmt.Println(i)
	close(ch)
}

