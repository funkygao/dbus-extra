package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	line := strings.Repeat("X", 100)
	if len(os.Args) > 1 {
		n, _ := strconv.Atoi(os.Args[1])
		line = strings.Repeat("X", n)
	}
	line += "\n"

	for {
		os.Stdout.WriteString(line)
	}

}
