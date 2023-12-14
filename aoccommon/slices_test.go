package aoccommon

import (
	"testing"
)

func TestTransposeEqualLength(t *testing.T) {
	input := []string{
		"abc",
		"fgh",
		"123",
	}

	runes2D := Map(input, func(a string) []rune {
		return []rune(a)
	})

	transposedRunes2D := Transpose(runes2D)

	actual := Map(transposedRunes2D, func(runes []rune) string {
		return string(runes)
	})

	expected := []string{
		"af1",
		"bg2",
		"ch3",
	}

	if len(actual) != len(expected) {
		t.Fail()
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestTransposeWide(t *testing.T) {
	input := []string{
		"abcde",
		"fghij",
		"12345",
	}

	runes2D := Map(input, func(a string) []rune {
		return []rune(a)
	})

	transposedRunes2D := Transpose(runes2D)

	actual := Map(transposedRunes2D, func(runes []rune) string {
		return string(runes)
	})

	expected := []string{
		"af1",
		"bg2",
		"ch3",
		"di4",
		"ej5",
	}

	if len(actual) != len(expected) {
		t.Fail()
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Fail()
		}
	}
}

func TestTransposeTall(t *testing.T) {
	input := []string{
		"af1",
		"bg2",
		"ch3",
		"di4",
		"ej5",
	}

	runes2D := Map(input, func(a string) []rune {
		return []rune(a)
	})

	transposedRunes2D := Transpose(runes2D)

	actual := Map(transposedRunes2D, func(runes []rune) string {
		return string(runes)
	})

	expected := []string{
		"abcde",
		"fghij",
		"12345",
	}

	if len(actual) != len(expected) {
		t.Fail()
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Fail()
		}
	}
}
