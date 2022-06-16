package main

func Match(pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	} else if len(text) == 0 {
		return false
	} else if pattern[0] == text[0] || pattern[0] == '.' {
		return Match(pattern[1:], text[1:])
	}
	return Match(pattern, text[1:])
}
