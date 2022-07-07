package iteration

import "testing"

func TestIteration(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, expected, got string) {
		t.Helper()
		if expected != got {
			t.Errorf("expected %q got %q", expected, got)
		}
	}

	t.Run("print 'a' 5 times", func(t *testing.T) {
		repeated := UseRepeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("print 'a' 6 times if count value given is 6", func(t *testing.T) {
		repeated := UseRepeat("a", 6)
		expected := "aaaaaa"
		assertCorrectMessage(t, expected, repeated)
	})

}

// * Benchmarking
// * go test -bench=.
// * go test -bench=<file-name.go>
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseRepeat("a", 5)
	}
}
