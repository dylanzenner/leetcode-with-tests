package sumDiff

import "testing"

func TestSumDiff(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := SumDiff(10, 3)
		want := 19

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := SumDiff(5, 6)
		want := 15

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}