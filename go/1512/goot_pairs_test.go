package goodPairs

import "testing"

func TestGoodPairs(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := GoodPairs([]int{1, 2, 3, 1, 1, 3})
		want := 4

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := GoodPairs([]int{1, 1, 1, 1})
		want := 6

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}