package hello10

import "testing"

func TestHello(t *testing.T) {
	if got := Hello(); got != "hello" {
		t.Errorf("got: %s", got)
	}
}
