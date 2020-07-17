package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"
)

func arbitraryMath(i int, c chan float64) {
	// Some stupid math to consume time
	f := 1.0 / float64(i)
	for j := 0; j < 10000; j++ {
		f = math.Sin(f) + 0.1
	}
	c <- f
}

const nbrRoutines = 100000

func main() {
	// NumCPU(): The set of available CPUs is checked by querying the operating system at process startup
	// GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously
	fmt.Println("Default configuration")
	fmt.Println("NumCPU =", runtime.NumCPU())
	fmt.Println("GOMAXPROCS =", runtime.GOMAXPROCS(0))
	fmt.Println()

	// Let the user set GOMAXPROCS via command line arguments
	if len(os.Args) == 2 {
		x, err := strconv.Atoi(os.Args[1])
		if err == nil {
			runtime.GOMAXPROCS(x)
			fmt.Println("Set GOMAXPROCS to", runtime.GOMAXPROCS(0))
			fmt.Println()
		} else {
			fmt.Println("First command line argument should be an integer")
			os.Exit(1)
		}
	}

	// Create channel, used to wait for all routines to complete
	c := make(chan float64)

	start := time.Now()
	for i := 0; i < nbrRoutines; i++ {
		// fmt.Println("Create routine ", i)
		go arbitraryMath(i, c)
	}
	fmt.Println("All routines created")

	for i := 0; i < nbrRoutines; i++ {
		x := <-c
		x = x + 1
		// fmt.Println("Recieved from routine", x-1)
	}
	duration := time.Since(start)
	fmt.Println("All routines done")
	fmt.Println("Execution time =", duration)
}
