package arrays

import "testing"

func TestSum(t *testing.T) {
	// * alternative delcaration
	// * ... means determined length automatically
	// * numbers := [...]int{1, 2, 3, 4, 5}
	numbers := [5]int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("expected %d got %d, given %v", expected, result, numbers)
	}
}
