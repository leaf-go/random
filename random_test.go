package random

import (
	"testing"
)

func TestD(t *testing.T) {
	min, max := 1000, 6666

	r := D(min, max)
	if r < min || r > max {
		t.Errorf("failed found: %v\n", r)
	}

	t.Logf("get n: %d", r)
}

func TestLenD(t *testing.T) {
	t.Logf("%d\n", LenD(3))
}

func TestW(t *testing.T) {
	t.Logf("%s\n", W(182))
}

func TestWs(t *testing.T) {
	s := Ws(128)
	t.Logf("ws:%s\n", s)
}
