package main

import (
	"fmt"
	"os"
)

// etl_toolbox - Extract transform load utilities
func etl_toolbox(path string) {
	fmt.Println("========================================")
	fmt.Println("  ETL-Toolbox")
	fmt.Println("  Extract transform load utilities")
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
	etl_toolbox(path)
}
