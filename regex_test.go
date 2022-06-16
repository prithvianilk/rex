package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestSingleCharacterValid(t *testing.T) {
	err := handleTest("a", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterInvalid(t *testing.T) {
	err := handleTest("a", "b", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterPatternInLongTextValid(t *testing.T) {
	err := handleTest("a", "aa", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterPatternInLongTextAtEndValid(t *testing.T) {
	err := handleTest("a", "ba", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterPatternInLongTextInvalid(t *testing.T) {
	err := handleTest("a", "bb", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func handleTest(pattern, text string, exists bool) error {
	regex := CreateRegex(pattern, text)
	if regex.Match() != exists {
		var errorMessage string
		if !exists {
			errorMessage = fmt.Sprintf("Pattern %s does not exist in %s", pattern, text)
		} else {
			errorMessage = fmt.Sprintf("Pattern %s exists in %s", pattern, text)
		}
		return errors.New(errorMessage)
	}
	return nil
}
