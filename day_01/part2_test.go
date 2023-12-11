package day01

import "testing"

func TestPart2(t *testing.T) {
	got := Part2("input2.txt")
	want := 281

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}
}
