package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := " ", " "
	for i, arg := range os.Args {
		s += seq + arg
		// seq = " "
		fmt.Printf("%d %s", i, s)
	}
}
