package mocking

import (
	"bytes"
	"reflect"
	"testing"
)

// * Spies are a kind of mock which can record how a dependency is used
// * They can record the arguments sent in, how many times it has been called, etc.
// type SpySleeper struct {
// 	calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.calls += 1
// }

// func TestCountdown(t *testing.T) {
// 	buffer := &bytes.Buffer{}
// 	sleeper := &SpySleeper{}

// 	Countdown(buffer, sleeper)

// 	result := buffer.String()
// 	expected := `3
// 2
// 1
// Go!`

// 	if result != expected {
// 		t.Errorf("expected %q result is %q", expected, result)
// 	}

// 	if sleeper.calls != 3 {
// 		t.Errorf("should sleep 3 times, slept only %d time(s)", sleeper.calls)
// 	}
// }

const write = "Write"
const sleep = "Sleep"

type SpyCountdownOperations struct {
	calls []string
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return
}

func (s *SpyCountdownOperations) Sleep() {
	s.calls = append(s.calls, sleep)
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		spy := &SpyCountdownOperations{}

		Countdown(buffer, spy)

		result := buffer.String()
		expectedPrint := `3
2
1
Go!`
		expectedCallsLength := 3

		if result != expectedPrint {
			t.Errorf("expected %q result is %q", expectedPrint, result)
		}

		if len(spy.calls) != expectedCallsLength {
			t.Errorf("should sleep %d times, slept only %d time(s)", expectedCallsLength, len(spy.calls))
		}

	})

	t.Run("sleeps before each print", func(t *testing.T) {
		spy := &SpyCountdownOperations{}

		Countdown(spy, spy)

		expected := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expected, spy.calls) {
			t.Errorf("expected %v result is %v", expected, spy.calls)
		}
	})
}
