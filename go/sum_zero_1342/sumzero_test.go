package sumzero

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumZero(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := SumZero(14)
		want := 6
		handleError(got, want, t)

	})

	t.Run("", func(t *testing.T){
		got := SumZero(8)
		want := 4
		handleError(got, want, t)
	})
}