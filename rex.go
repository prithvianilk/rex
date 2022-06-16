package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pattern, text := os.Args[1], os.Args[2]
	lines := strings.Split(text, "\n")
	for lnum, line := range lines {
		if Match(pattern, line) {
			fmt.Println(lnum, line)
		}
	}
}
