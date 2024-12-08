package main

import (
	_ "embed"
	"testing"

	"github.com/limwa/advent-of-code-2024/lib/util"
)

//go:embed example.txt
var _example string

func init() {
	util.NormalizeInput(&_example)
}

func createTest(t *testing.T, solver func(string) string, expected string) {
	actual := solver(_example)
	util.NormalizeInput(&expected)

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}

func TestSolvePart1(t *testing.T) {
	expected := "11"
	createTest(t, solvePart1, expected)
}

func TestSolvePart2(t *testing.T) {
	expected := "31"
	createTest(t, solvePart2, expected)
}
