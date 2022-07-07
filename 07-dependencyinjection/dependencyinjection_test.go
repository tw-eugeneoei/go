package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	result := buffer.String()
	expected := "Hello, Chris"

	if expected != result {
		t.Errorf("expected %q got %q", expected, result)
	}
}
