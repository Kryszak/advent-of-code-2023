package day06

import "testing"

func TestPart2(t *testing.T) {
	got := Part2("input.txt")
	want := 71504

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
