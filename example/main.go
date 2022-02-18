package main

import (
	"fmt"
	"time"

	"github.com/mpragliola/stopwatch"
)

func main() {
	s := stopwatch.New()

	time.Sleep(500 * time.Millisecond)

	s.Mark("foo")

	time.Sleep(300 * time.Millisecond)

	s.Dump()

	fmt.Println(s.Json())
}
