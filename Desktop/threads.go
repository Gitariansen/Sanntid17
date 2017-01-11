package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int = 0

func thread_1_routine() {
	for x := 0; x <= 1000000; x++ {
		i++
	}
}

func thread_2_routine() {
	for x := 0; x <= 1000000; x++ {
		i--
	}
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	go thread_1_routine()
	go thread_2_routine()

	//time.Sleep(10*time.Millisecond)
	fmt.Println(i)
}

