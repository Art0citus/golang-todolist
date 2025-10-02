package main

import (
	"fmt"
	"os"
)

func main() {
	// saare CLI arguments ek slice me milenge
	args := os.Args

	// slice ko print kar ke dekh lo
	fmt.Println(args)

	// Example access
	if len(args) > 1 {
		fmt.Println("First argument:", args[1])
	}
}
