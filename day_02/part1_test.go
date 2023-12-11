package day02

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("input.txt")
	want := 8

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}