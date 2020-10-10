package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}

	end := time.Since(start).Seconds()
	fmt.Println(end)
	start = time.Now()

	strings.Join(os.Args[1:], "")
	fmt.Println(time.Since(start).Seconds())

}
