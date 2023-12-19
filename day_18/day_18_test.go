package day18

import "testing"

func TestDay18(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input.txt")
		want := 62

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input.txt")
		want := 952408144115

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
