package main

import (
	"fmt"
	"math"
	"time"
)

func arbInc(i int, c chan float64) {
	// Some stupid math to consume time
	f := 1.0 / float64(i)
	for j := 0; j < 10000; j++ {
		f = math.Sin(f) + 0.1
	}
	c <- f
}

const nbrRoutines = 100000

func main() {

	c := make(chan float64)

	start := time.Now()
	for i := 0; i < nbrRoutines; i++ {
		// fmt.Println("Create routine ", i)
		go arbInc(i, c)
	}
	fmt.Println("All routines created")

	for i := 0; i < nbrRoutines; i++ {
		x := <-c
		x = x + 1
		// fmt.Println("Recieved from routine", x-1)
	}
	duration := time.Since(start)
	fmt.Println("All routines done")
	fmt.Println("Duration =", duration)
	time.Sleep(time.Second)
	fmt.Println("Done")
}
