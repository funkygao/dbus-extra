// +build OMIT

package main

import (
	"fmt"
	"strings"
	"testing"
)

var chanSize int

func main() {
	chanSizes := []int{0, 1, 50, 100, 500, 1000, 5000}
	fmt.Println("benchmark channel of int")
	fmt.Println("chan buf size")
	for _, sz := range chanSizes {
		chanSize = sz
		fmt.Printf("%13d %s\n", chanSize, testing.Benchmark(benchmarkChanInt).String())
	}
	fmt.Println("benchmark channel of [100]byte")
	fmt.Println("chan buf size")
	for _, s := range chanSizes {
		chanSize = s
		fmt.Printf("%13d %s\n", chanSize, testing.Benchmark(benchmarkChanByte100).String())
	}
}

func benchmarkChanInt(b *testing.B) {
	ch := make(chan int, chanSize)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}

		close(ch)
	}()

	for _ = range ch {
	}
} //

func benchmarkChanByte100(b *testing.B) {
	ch := make(chan []byte, chanSize)
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
