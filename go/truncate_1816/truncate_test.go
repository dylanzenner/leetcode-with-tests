package truncate

import "testing"

func handleError(got, want string, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTruncate(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Truncate("the quick brown fox jumped over the lazy dog", 4)
		want := "the quick brown fox"
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := Truncate("one piece is awesome", 2)
		want := "one piece"
		handleError(got, want, t)
	})
}