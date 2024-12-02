package math_test

import (
	"testing"

	"github.com/limwa/advent-of-code-2024/lib/math"
)

func TestAbs(t *testing.T) {
	actual := math.Abs(123)
	expected := 123

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestAbsNegative(t *testing.T) {
	actual := math.Abs(-123)
	expected := 123

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSign(t *testing.T) {
	actual := math.Sign(123)
	expected := 1

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSignNegative(t *testing.T) {
	actual := math.Sign(-123)
	expected := -1

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSignZero(t *testing.T) {
	actual := math.Sign(0)
	expected := 0

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
