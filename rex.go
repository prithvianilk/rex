package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	pattern, text := os.Args[1], os.Args[2]
	lines := strings.Split(text, "\n")
	red := color.New(color.FgRed)

	for _, line := range lines {
		match, err := Match(pattern, line)
		if err != nil {
			continue
		}
		matchIndex, err := getMatchIndex(line, match)
		if err != nil {
			panic(err)
		}
		for i := range line {
			if i >= matchIndex && i < matchIndex+len(match) {
				red.Print(string(line[i]))
			} else {
				fmt.Print(string(line[i]))
			}
		}
		fmt.Println()
	}
}
