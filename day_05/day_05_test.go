package day05

import "testing"

func TestDay05(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input.txt")
		want := 35

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input.txt")
		want := 46

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
