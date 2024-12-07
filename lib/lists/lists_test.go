package lists_test

import (
	"testing"

	"github.com/limwa/advent-of-code-2024/lib/lists"
)

func TestSumEmpty(t *testing.T) {
	actual := lists.Sum([]int{})
	expected := 0

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSumSingle(t *testing.T) {
	actual := lists.Sum([]int{1})
	expected := 1

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSumMultiple(t *testing.T) {
	actual := lists.Sum([]int{1, 5, 3})
	expected := 9

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestContainsEmpty(t *testing.T) {
	actual := lists.Contains([]int{}, 1)
	expected := false

	if actual != expected {
		t.Errorf("expected %t, got %t", expected, actual)
	}
}

func TestContainsSingle(t *testing.T) {
	actual := lists.Contains([]int{1}, 1)
	expected := true

	if actual != expected {
		t.Errorf("expected %t, got %t", expected, actual)
	}
}

func TestContainsMultiple(t *testing.T) {
	actual := lists.Contains([]int{1, 5, 3}, 3)
	expected := true

	if actual != expected {
		t.Errorf("expected %t, got %t", expected, actual)
	}
}
