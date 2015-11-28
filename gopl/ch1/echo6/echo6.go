// Echo6 measure performance between Strings.join and inefficient loop
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func timeElapsed(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}

func stringsJoin() {
	defer timeElapsed(time.Now(), "stringsJoin")

	fmt.Println(strings.Join(os.Args[1:], " "))
}

func slowStringJoin() {
	defer timeElapsed(time.Now(), "slowStringJoin")

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	stringsJoin()
	slowStringJoin()
}
