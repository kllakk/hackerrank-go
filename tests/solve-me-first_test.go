package tests

import (
	solveMeFirst "hackerrank-go/src/solve-me-first"
	"testing"
)

func TestSolveMeFirst01(t *testing.T) {
	var v uint32

	v = solveMeFirst.SolveMeFirst(1, 2)

	if v != 3 {
		t.Error("Expected 3, got ", v)
	}
}