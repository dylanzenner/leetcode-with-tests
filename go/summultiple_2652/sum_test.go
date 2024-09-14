package sumMultiple

import "testing"

func TestSumMultiple(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := SumMultiple(7)
		want := 21

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := SumMultiple(10)
		want := 40

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}