package aoccommon

import (
	"strings"
	"testing"
)

func TestSortedQueue_InsertSimple(t *testing.T) {
	input := []int{1, 5, 7, 2, 8, 4, 2}
	sortedQueue := NewSortedQueue[int](func(a, b int) int {
		return a - b
	})

	for _, in := range input {
		sortedQueue.Insert(in)
	}

	expected := []int{1, 2, 2, 4, 5, 7, 8}
	for _, expectedValue := range expected {
		actualValue, exists := sortedQueue.Pop()
		if !exists {
			t.Errorf("queue contains not enough values")
		}

		if actualValue != expectedValue {
			t.Errorf("queue did not sort the values as expected: got %v instead of %v", actualValue, expectedValue)
		}
	}
}

func TestSortedQueue_InsertComplex(t *testing.T) {
	type labeledInt struct {
		val   int
		label string
	}

	input := []labeledInt{
		{1, "hello"},
		{5, "world"},
		{7, "this"},
		{2, "is"},
		{8, "a"},
		{4, "test"},
		{2, "!"},
	}
	sortedQueue := NewSortedQueue[labeledInt](func(a, b labeledInt) int {
		if a.val-b.val == 0 {
			return strings.Compare(a.label, b.label)
		}
		return a.val - b.val
	})

	for _, in := range input {
		sortedQueue.Insert(in)
	}

	expected := []labeledInt{
		{1, "hello"},
		{2, "!"},
		{2, "is"},
		{4, "test"},
		{5, "world"},
		{7, "this"},
		{8, "a"},
	}
	for _, expectedValue := range expected {
		actualValue, exists := sortedQueue.Pop()
		if !exists {
			t.Errorf("queue contains not enough values")
		}

		if actualValue != expectedValue {
			t.Errorf("queue did not sort the values as expected: got %v instead of %v", actualValue, expectedValue)
		}
	}
}

func TestPosition2D_InBounds(t *testing.T) {
	pos := Position2D{
		X: 4,
		Y: 5,
	}

	if pos.InBounds(6, 3) {
		t.Errorf("position is out of bounds, Y exceeds the maximum height")
	}
	if pos.InBounds(4, 6) {
		t.Errorf("position is out of bounds, X exceeds the maximum width")
	}

	if !pos.InBounds(6, 7) {
		t.Errorf("position is incorrectly detected as out of bounds")
	}

	pos = Position2D{
		X: 0,
		Y: -1,
	}
	if pos.InBounds(5, 5) {
		t.Errorf("position is out of bounds, Y is below 0")
	}
}
