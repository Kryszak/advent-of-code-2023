package day15

import "testing"

func TestPart2(t *testing.T) {
	got := Part2("input.txt")
	want := 145

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}