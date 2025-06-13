package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Task 1.1, 1.2, 1.3
func main() {
	start := time.Now()
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println("index: " + strconv.Itoa(i) + " value: " + arg)
	}
	secs := time.Since(start).Seconds()

	elapsedstr := strconv.FormatFloat(secs, 'E', -1, 64)

	fmt.Println("time elapsed: " + elapsedstr)
}
