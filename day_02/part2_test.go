package day02

import "testing"

func TestPart2(t *testing.T) {
	got := Part2("input.txt")
	want := 2286

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
