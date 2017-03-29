package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/funkygao/golib/gofmt"
	"github.com/funkygao/golib/io"
)

var n int64

func main() {
	go func() {
		var lastN int64
		for {
			time.Sleep(time.Second * 2)
			fmt.Printf("%s/s\n", gofmt.Comma((n-lastN)/2))

			lastN = n
		}
	}()
	r := bufio.NewReader(os.Stdin)
	for {
		io.ReadLine(r)
		n++
	}
}
