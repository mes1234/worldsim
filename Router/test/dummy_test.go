package dummy_test

import (
	"testing"

	"github.com/mes1234/worldsim/Router/internal/dummy"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := dummy.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
