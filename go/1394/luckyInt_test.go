package luckyInt

import "testing"

func TestLuckyInt(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := LuckyInt([]int{2, 2, 3, 4})
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := LuckyInt([]int{1, 2, 2, 3, 3, 3})
		want := 3

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 3", func(t *testing.T){
		got := LuckyInt([]int{2, 2, 2, 3, 3})
		want := -1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}