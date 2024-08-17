package maxones

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxOnes(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := MaxOnes([]int{1, 0, 1, 1, 1, 0, 0, 1})
		want := 3

		handleError(got, want, t)
	})

	t.Run("no ones", func(t *testing.T){
		got := MaxOnes([]int{0, 0, 0, 0})
		want := 0

		handleError(got, want, t)
	})

	t.Run("all ones", func(t *testing.T){
		got := MaxOnes([]int{1, 1, 1})
		want := 3
		handleError(got, want, t)
	})
}