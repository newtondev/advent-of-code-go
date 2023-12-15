package main

import "testing"

func TestSolvePartOne(t *testing.T) {
	expected := uint32(142)
	result := solvePartOne(readFileInput("../../test_inputs/01_one.txt"))
	if expected != result {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestSolvePartTwo(t *testing.T) {
	expected := uint32(281)
	result := solvePartTwo(readFileInput("../../test_inputs/01_two.txt"))
	if expected != result {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
