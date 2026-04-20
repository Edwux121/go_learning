package depInjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Edvinas")

	got := buffer.String()
	want := "Hello, Edvinas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
