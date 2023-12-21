package day20

import "testing"

func TestDay20(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input.txt")
		want := 11687500

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
