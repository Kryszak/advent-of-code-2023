package day02

import "testing"

func TestDay02(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input.txt")
		want := 8

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input.txt")
		want := 2286

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
