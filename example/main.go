package main

import (
	"fmt"
	"time"

	"github.com/mpragliola/stopwatch"
)

func main() {
	s := stopwatch.NewStart()

	time.Sleep(500 * time.Millisecond)

	s.Mark("foo")

	time.Sleep(300 * time.Millisecond)

	fmt.Print("\nData object dump\n")
	fmt.Println(s.Data())
	fmt.Print("\nDirect dump\n")
	s.Dump()
	fmt.Print("\nJson dump\n")
	fmt.Println(s.Json())
}
