package day08

import "testing"

func TestDay08(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("input1.txt")
		want := 2

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2("input2.txt")
		want := 6

		if got != want {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
