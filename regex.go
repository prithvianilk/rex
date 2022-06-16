package main

func Match(pattern, text string) bool {
	if matchFromStart(pattern, text) {
		return true
	} else if len(text) > 0 {
		return Match(pattern, text[1:])
	}
	return false
}

func matchFromStart(pattern, text string) bool {
	if len(pattern) == 0 {
		return true
	} else if len(pattern) > 1 && pattern[1] == '*' {
		return matchStar(pattern[0], pattern[2:], text)
	} else if len(text) == 0 {
		return false
	} else if pattern[0] == text[0] || pattern[0] == '.' {
		return matchFromStart(pattern[1:], text[1:])
	}
	return false
}

func matchStar(char byte, patten, text string) bool {
	acceptsAny := char == '.'
	for !matchFromStart(patten, text) {
		if len(text) > 0 && (acceptsAny || char == text[0]) {
			text = text[1:]
		} else {
			return false
		}
	}
	return true
}
