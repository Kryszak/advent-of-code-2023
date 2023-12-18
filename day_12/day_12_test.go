package day12

import "testing"

func TestDay12(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input.txt")
		want := 21

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input.txt")
		want := 525152

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
