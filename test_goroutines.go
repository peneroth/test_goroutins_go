package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("go routine thread")

	for i := 0; i < 3; i++ {
		go func(msg string) {
			fmt.Println(msg)
		}("Main thread")
	}

	time.Sleep(time.Second)
	fmt.Println("done")
}
