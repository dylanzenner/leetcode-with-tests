package maxProd

import "testing"

func TestMaxProd(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := MaxProd([]int{3, 4, 5, 2})
		want := 12

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := MaxProd([]int{1, 5, 4, 5})
		want := 16

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}