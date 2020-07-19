package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"time"
)

const nbrGoRoutines = 100000
const nbrLoops = 10000

func arbitraryMath(i int, loops int, c chan float64) {
	// Some stupid math to consume time
	f := 1.0 / float64(i)
	for j := 0; j < loops; j++ {
		f = math.Sin(f) + 0.1
	}
	c <- f
}

func main() {
	// initiate
	loops := nbrLoops           // default number of loops
	goRoutines := nbrGoRoutines // default nbr of go routines
	//
	// parse command line arguments
	pos := 1
	for pos < len(os.Args) {
		switch os.Args[pos] {
		case "-help":
			fmt.Println("Syntax: test_goroutines [argument [parameter]]")
			fmt.Println("-help, this text")
			fmt.Println("-threads x, maximum number of OS threads")
			fmt.Println("-goroutines x, number of routines to be created")
			fmt.Println("-loops x, number of loops in arbiraryMath")
			fmt.Println()
			fmt.Println("Default configuration")
			fmt.Println("NumCPU =", runtime.NumCPU())
			fmt.Println("GOMAXPROCS =", runtime.GOMAXPROCS(0))
			fmt.Println("goroutines = ", nbrGoRoutines)
			fmt.Println("loops = ", nbrLoops)
			os.Exit(0)
		case "-threads":
			pos++
			if pos >= len(os.Args) {
				fmt.Println("-threads requires a paramter")
				os.Exit(1)
			}
			x, err := strconv.Atoi(os.Args[pos])
			if err == nil {
				runtime.GOMAXPROCS(x)
				fmt.Println("Set GOMAXPROCS to", runtime.GOMAXPROCS(0))
			} else {
				fmt.Println("-threads requires a paramter of type integer")
				os.Exit(1)
			}
			pos++
		case "-goroutines":
			pos++
			if pos >= len(os.Args) {
				fmt.Println("-goroutines requires a paramter")
				os.Exit(1)
			}
			x, err := strconv.Atoi(os.Args[pos])
			if err == nil {
				goRoutines = x
				fmt.Println("Set number of go routines to ", goRoutines)
			} else {
				fmt.Println("-goroutines requires a paramter of type integer")
				os.Exit(1)
			}
			pos++
		case "-loops":
			pos++
			if pos >= len(os.Args) {
				fmt.Println("-loops requires a paramter")
				os.Exit(1)
			}
			x, err := strconv.Atoi(os.Args[pos])
			if err == nil {
				loops = x
				fmt.Println("Set number of loops to ", x)
			} else {
				fmt.Println("-loops requires a paramter of type integer")
				os.Exit(1)
			}
			pos++
		default:
			fmt.Println(os.Args[pos], "is an unrecognized command")
			os.Exit(1)
		}
	}

	// Create channel, used to wait for all routines to complete
	c := make(chan float64)

	start := time.Now()
	for i := 0; i < goRoutines; i++ {
		// fmt.Println("Create routine ", i)
		go arbitraryMath(i, loops, c)
	}
	fmt.Println("All routines created")

	for i := 0; i < goRoutines; i++ {
		x := <-c
		x = x + 1
		// fmt.Println("Recieved from routine", x-1)
	}
	duration := time.Since(start)
	fmt.Println("All routines done")
	fmt.Println("Execution time =", duration)
}
