package day16

import "testing"

func TestPart2(t *testing.T) {
	got := Part2("input.txt")
	want := 51 

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
