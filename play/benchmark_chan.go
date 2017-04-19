package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/funkygao/gobench/util"
)

var buf int

func main() {
	bs := []int{0, 1, 2, 5, 10, 50, 100, 200, 500, 1000, 5000}
	fmt.Printf("%5s %s\n", "qlen", "with int")
	for _, s := range bs {
		buf = s
		fmt.Printf("%5d %s\n", buf, testing.Benchmark(benchmarkChanInt).String())
	}
	fmt.Printf("%5s %s\n", "qlen", "with [100]byte")
	for _, s := range bs {
		buf = s
		fmt.Printf("%5d %s\n", buf, testing.Benchmark(benchmarkChanByte100).String())
	}
}

// START OMIT
func benchmarkChanInt(b *testing.B) {
	ch := make(chan int, buf)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}

		close(ch)
	}()

	for _ = range ch {
	}
}

// END OMIT

func benchmarkChanByte100(b *testing.B) {
	ch := make(chan []byte, buf)
	data := []byte(strings.Repeat("x", 100))
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- data
		}

		close(ch)
	}()

	for _ = range ch {
	}
}
