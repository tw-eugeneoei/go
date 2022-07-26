package synctut

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("value should be 3 after incrementing 3 times", func(t *testing.T) {
		expected := 3

		counter := Counter{}
		counter.inc()
		counter.inc()
		counter.inc()

		assertCounter(t, expected, &counter)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		// * A "WaitGroup" waits for a collection of goroutines to finish
		var wg sync.WaitGroup
		// * main goroutine calls "Add" to set the number of goroutines to wait for.
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.inc()
				wg.Done()
			}()
		}

		// * "Wait" can be used to block until all goroutines have finished
		wg.Wait()

		assertCounter(t, wantedCount, &counter)
	})
}

func assertCounter(t testing.TB, expected int, counter *Counter) {
	t.Helper()
	got := counter.value()
	if expected != got {
		t.Errorf("expected %d got %d", expected, got)
	}
}
