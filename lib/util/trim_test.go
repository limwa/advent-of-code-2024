package util_test

import (
	"testing"

	"github.com/limwa/advent-of-code-2024/lib/util"
)

func TestEmptyInput(t *testing.T) {
	input := "\n\r\n"

	util.NormalizeInput(&input)

	if input != "" {
		t.Errorf("expected empty string, got %q", input)
	}
}

func TestCarriageReturnInput(t *testing.T) {
	input := "hi\r\n"
	expected := "hi"

	util.NormalizeInput(&input)

	if input != expected {
		t.Errorf("expected %q, got %q", expected, input)
	}
}

func TestMultilineInput(t *testing.T) {
	input := "hi\n\r\nthere\n\n"
	expected := "hi\n\nthere"

	util.NormalizeInput(&input)

	if input != expected {
		t.Errorf("expected %q, got %q", expected, input)
	}
}
