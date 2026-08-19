package main

import (
	"fmt"
	"os"
)

// version_control_helpers - Git workflow automation
func version_control_helpers(path string) {
	fmt.Println("========================================")
	fmt.Println("  Version-Control-Helpers")
	fmt.Println("  Git workflow automation")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	version_control_helpers(path)
}
