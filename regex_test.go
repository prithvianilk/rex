package main

import (
	"fmt"
	"testing"
)

func TestSingleCharacterValid(t *testing.T) {
	err := handleTest("a", "a", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterInvalid(t *testing.T) {
	err := handleTest("a", "b", "", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterPatternInLongTextValid(t *testing.T) {
	err := handleTest("a", "aa", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterPatternInLongTextAtEndValid(t *testing.T) {
	err := handleTest("a", "ba", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestLongPatternInLongTextInMiddleValid(t *testing.T) {
	err := handleTest("12", "333321233", "12", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestLongPatternInLongTextInvalid(t *testing.T) {
	err := handleTest("12", "21321", "", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSingleCharacterPatternInLongTextInvalid(t *testing.T) {
	err := handleTest("a", "bb", "", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDotPatternValid(t *testing.T) {
	err := handleTest(".", "a", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDotPatternInLongTextValid(t *testing.T) {
	err := handleTest(".", "aa", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestDotPatternInNoText(t *testing.T) {
	err := handleTest(".", "", "", false)
	if err != nil {
		t.Error(err.Error())
	}
}
func TestStarPattern(t *testing.T) {
	err := handleTest("a.*", "a", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternLong(t *testing.T) {
	err := handleTest("a.*", "abc", "a", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternNothing(t *testing.T) {
	err := handleTest(".*", "", "", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternSameCharacters(t *testing.T) {
	err := handleTest("ba*", "baaa", "b", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternDifferentCharacters(t *testing.T) {
	err := handleTest("ab*c", "abbbc", "abbbc", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternWithWordEndsWithValid(t *testing.T) {
	err := handleTest(".*a", "bcda", "bcda", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternWithWordEndsWithInvalid(t *testing.T) {
	err := handleTest(".*aa", "bba", "", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternComplexInMiddleValid(t *testing.T) {
	err := handleTest(".*a.*", "bcabc", "bca", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternComplexInMiddleInvalid(t *testing.T) {
	err := handleTest(".*a.*", "bcdbc", "", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternComplexInMiddleBigValid(t *testing.T) {
	err := handleTest(".*a.*", "bcaabc", "bca", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternComplexInMiddleBigInvalid(t *testing.T) {
	err := handleTest(".*aa.*", "bcabc", "", false)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestStarPatternComplexInMiddleWithSurrounding(t *testing.T) {
	err := handleTest("0.*a.*1", "0bcaabc1", "0bcaabc1", true)
	if err != nil {
		t.Error(err.Error())
	}
}

func handleTest(pattern, text string, expectedMatch string, exists bool) error {
	match, err := Match(pattern, text)
	if (err == nil) != exists {
		if !exists {
			return fmt.Errorf("Pattern %s does not exist in %s", pattern, text)
		}
		return fmt.Errorf("Pattern %s exists in %s", pattern, text)
	}
	if expectedMatch != match {
		return fmt.Errorf("Expected regex match to be %s, but got %s", expectedMatch, match)
	}
	return nil
}
