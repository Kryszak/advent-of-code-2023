package day03

import "testing"

func TestDay03(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input.txt")
		want := 4361

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input.txt")
		want := 467835

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})
}
