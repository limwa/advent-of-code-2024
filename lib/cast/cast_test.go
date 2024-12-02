package cast_test

import (
	"testing"
	"reflect"

	"github.com/limwa/advent-of-code-2024/lib/cast"
)

func TestToInt(t *testing.T) {
	actual := cast.ToInt("123")
	expected := 123

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestToIntError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, got none")
		}
	}()

	cast.ToInt("abc")
}

func TestToIntSlice(t *testing.T) {
	actual := cast.ToIntSlice([]string{"1", "2", "3"})
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
