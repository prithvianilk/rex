package main

type Regex struct {
	pattern string
	text    string
}

func CreateRegex(pattern, text string) *Regex {
	return &Regex{pattern: pattern, text: text}
}

func (r *Regex) Match() bool {
	return true
}
