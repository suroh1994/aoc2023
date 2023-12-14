package aoccommon

import (
	"slices"
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

func TestRotateRightEqual(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	expected := [][]int{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}

	actual := RotateRight(input)

	if len(actual) != len(expected) {
		t.Fail()
	}

	for i := range expected {
		if !slices.Equal(actual[i], expected[i]) {
			t.Fail()
		}
	}
}

func TestRotateRightTall(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{5, 6, 7},
		{9, 10, 11},
		{13, 14, 15},
	}

	expected := [][]int{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
	}

	actual := RotateRight(input)

	if len(actual) != len(expected) {
		t.Fail()
	}

	for i := range expected {
		if !slices.Equal(actual[i], expected[i]) {
			t.Fail()
		}
	}
}

func TestRotateRightWide(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	expected := [][]int{
		{9, 5, 1},
		{10, 6, 2},
		{11, 7, 3},
		{12, 8, 4},
	}

	actual := RotateRight(input)

	if len(actual) != len(expected) {
		t.Fail()
	}

	for i := range expected {
		if !slices.Equal(actual[i], expected[i]) {
			t.Fail()
		}
	}
}
