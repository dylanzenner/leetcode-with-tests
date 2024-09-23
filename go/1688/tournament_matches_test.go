package matches

import "testing"

func TestNumMatches(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := NumMatches(7)
		want := 6

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := NumMatches(14)
		want := 13

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}