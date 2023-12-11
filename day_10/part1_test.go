package day10

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("input1.txt")
	want := 4

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
