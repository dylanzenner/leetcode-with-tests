package maxnum

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMaxNum(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := MaxNum(3, 2)
		want := 7
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := MaxNum(4, 1)
		want := 6
		handleError(got, want, t)
	})
}