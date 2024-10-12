package heightChecker

import "testing"

func TestHeightChecker(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := HeightChecker([]int{1, 1, 4, 2, 1, 3})
		want := 3

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := HeightChecker([]int{5, 1, 2, 3, 4})
		want := 5

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}