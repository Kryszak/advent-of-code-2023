package day03

import "testing"

func TestPart2(t *testing.T) {
	got := Part2("input.txt")
	want := 467835

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}