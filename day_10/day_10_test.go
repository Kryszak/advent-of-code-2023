package day10

import "testing"

func TestDay10(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input1.txt")
		want := 4

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input2.txt")
		want := 4

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
