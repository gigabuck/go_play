// Echo4 prints index and arg, one per line
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%d - %s\n", i, os.Args[i])
	}
}
