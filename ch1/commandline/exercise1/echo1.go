package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	sep := " "
	for _, arg := range os.Args {
		s += arg + sep
	}
	fmt.Println(s)
}
