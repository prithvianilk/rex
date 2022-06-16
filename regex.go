package main

type Regex struct {
	pattern string
	text    string
}

func CreateRegex(pattern, text string) *Regex {
	return &Regex{pattern: pattern, text: text}
}

func (r *Regex) Match() bool {
	if len(r.pattern) == 0 {
		return true
	} else if len(r.text) == 0 {
		return false
	} else if r.pattern[0] == r.text[0] || r.pattern[0] == '.' {
		r.pattern, r.text = r.pattern[1:], r.text[1:]
		return r.Match()
	}
	r.text = r.text[1:]
	return r.Match()
}
