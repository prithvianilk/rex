package main

import "errors"

func invalidError() (string, error) {
	return "", errors.New("no match found for given pattern")
}

func Match(pattern, text string) (string, error) {
	match, err := matchFromStart(pattern, text)
	if err == nil {
		return match, nil
	} else if len(text) > 0 {
		return Match(pattern, text[1:])
	}
	return invalidError()
}

func matchFromStart(pattern, text string) (string, error) {
	if len(pattern) == 0 {
		return "", nil
	} else if len(pattern) > 1 && pattern[1] == '*' {
		return matchStar(pattern[0], pattern[2:], text)
	} else if len(text) > 0 && (pattern[0] == text[0] || pattern[0] == '.') {
		match, err := matchFromStart(pattern[1:], text[1:])
		if err == nil {
			return string(text[0]) + match, nil
		}
	}
	return invalidError()
}

func matchStar(char byte, patten, text string) (string, error) {
	acceptsAny := char == '.'
	prefix := ""
	for {
		match, err := matchFromStart(patten, text)
		if err == nil {
			return prefix + match, nil
		}
		if len(text) > 0 && (acceptsAny || char == text[0]) {
			prefix += string(text[0])
			text = text[1:]
			continue
		}
		return invalidError()
	}
}
