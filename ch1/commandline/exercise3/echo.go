package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	loop()
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("Loop took %dms\n", elapsed)

	start = time.Now()
	join()
	t = time.Now()
	elapsed = t.Sub(start)

	fmt.Printf("Join took %dms\n", elapsed)
}

func loop() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func join() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
