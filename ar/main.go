package main

import (
	"fmt"
	"os"
)

func main() {
	// Print the original command-line arguments
	fmt.Printf("Original Args: %#v\n", os.Args)

	// Create a copy of the command-line arguments
	newArgs := make([]string, len(os.Args))
	copy(newArgs, os.Args)

	// Modify the first argument
	newArgs[0] = "newExecutablePath"

	// Print the modified arguments
	fmt.Printf("Modified Args: %#v\n", newArgs)
}
