package day06

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("input.txt")
	want := 288

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
