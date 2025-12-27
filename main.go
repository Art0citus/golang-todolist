package main

import (
	"fmt"
	"os"
)

func main() {
	
	args := os.Args

	fmt.Println(args)

	
	if len(args) > 1 {
		fmt.Println("First argument:", args[1])
	}
}
