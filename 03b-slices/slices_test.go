package slices

import (
	"reflect"
	"testing"
)

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	result := SumSlice(numbers)
	expected := 15

	if result != expected {
		t.Errorf("expected %d got %d, given %v", expected, result, numbers)
	}
}

func TestSumAll(t *testing.T) {

	sliceOne := []int{1, 2, 3}
	sliceTwo := []int{10, 20, 30}

	result := SumAll(sliceOne, sliceTwo)
	expected := []int{6, 60}

	// * Go does not let you use equality operators with slices
	// * ie result != expected is invalid
	// if result != expected {
	// 	t.Errorf("expected %d got %d, given %v and %v", expected, result, sliceOne, sliceTwo)
	// }

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %d got %d, given %v and %v", expected, result, sliceOne, sliceTwo)
	}
}

func TestSumAllTails(t *testing.T) {

	// * A handy side-effect of this is this adds a little type-safety to our code.
	// * If a developer mistakenly adds a new test with checkSums(t, got, "dave") the compiler will stop them in their tracks.
	checkSums := func(t testing.TB, expected, result, arrayOne, arrayTwo []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %d got %d, given %v and %v", expected, result, arrayOne, arrayTwo)
		}
	}

	t.Run("should return [5, 50]", func(t *testing.T) {
		sliceOne := []int{1, 2, 3}
		sliceTwo := []int{10, 20, 30}

		result := SumAllTails(sliceOne, sliceTwo)
		expected := []int{5, 50}
		checkSums(t, expected, result, sliceOne, sliceTwo)
	})

	t.Run("should return [0, 50]", func(t *testing.T) {
		sliceOne := []int{}
		sliceTwo := []int{10, 20, 30}

		result := SumAllTails(sliceOne, sliceTwo)
		expected := []int{0, 50}
		checkSums(t, expected, result, sliceOne, sliceTwo)
	})

}
