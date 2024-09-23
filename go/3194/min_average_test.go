package minAverage

import "testing"

func TestMinAverage(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := MinAverage([]int{7, 8, 3, 4, 15, 13, 4, 1})
		want := 5.5

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := MinAverage([]int{1, 9, 8, 3, 10, 5})
		want := 5.5

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	
}