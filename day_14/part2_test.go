package day14

import "testing"

func TestPart1(t *testing.T) {
	got := Part2("input.txt")
	want := 0

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
