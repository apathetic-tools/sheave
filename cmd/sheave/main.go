package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define the base flags
	versionFlag := flag.Bool("version", false, "Print version information")

	// Parse the command line arguments
	flag.Parse()

	if *versionFlag {
		fmt.Println("sheave version 0.2.0")
		os.Exit(0)
	}

	fmt.Println("Hello world")
}
