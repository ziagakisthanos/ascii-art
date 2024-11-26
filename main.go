package main

import (
	"ascii/ascii"
	"fmt"
	"os"
)

func main() {
	// Process command-line arguments and handle errors
	err := ascii.HandleArgs(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
