package common

import (
	"strings"
	"testing"
)

func TestHash(t *testing.T) {
	input := strings.Split("HASH,rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ",")
	expected := []int{52, 30, 253, 97, 47, 14, 180, 9, 197, 48, 214, 231}

	var actual []int
	for _, in := range input {
		actual = append(actual, Hash(in))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("incorrect hash: expected %d but got %d", expected[i], actual[i])
		}
	}
}
