package main

import (
	"fmt"
)

func getMatchIndex(text, pattern string) (int, error) {
	n, m := len(text), len(pattern)
	for i := 0; i+m-1 < n; i++ {
		isMatching := true
		for j := range pattern {
			if text[i+j] != pattern[j] {
				isMatching = false
				break
			}
		}
		if isMatching {
			return i, nil
		}
	}
	return 0, fmt.Errorf("pattern %s not found in text %s", pattern, text)
}
