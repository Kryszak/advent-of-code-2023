package day19

import "testing"

func TestDay19(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input.txt")
		want := 19114

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input.txt")
		want := 167409079868000

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
