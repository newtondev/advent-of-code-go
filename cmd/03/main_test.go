package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	expected := uint32(4361)
	result := solve(readFileInput("../../test_inputs/03_one.txt"))
	if expected != result.p1 {
		t.Fatalf("expected %d, got %d", expected, result.p1)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := uint32(467835)
	result := solve(readFileInput("../../test_inputs/03_one.txt"))
	if expected != result.p2 {
		t.Fatalf("expected %d, got %d", expected, result.p2)
	}
}
