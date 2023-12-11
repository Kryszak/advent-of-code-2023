package day03

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("input.txt")
	want := 4361

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
