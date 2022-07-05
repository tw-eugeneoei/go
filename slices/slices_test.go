package slices

import "testing"

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	result := SumSlice(numbers)
	expected := 15

	if result != expected {
		t.Errorf("expected %d got %d, given %v", expected, result, numbers)
	}
}
