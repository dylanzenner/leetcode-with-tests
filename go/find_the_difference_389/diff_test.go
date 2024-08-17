package diff

import "testing"

func handleError(got, want byte, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDiff(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Diff("apple", "apples")
		want := byte('s')
		handleError(got, want, t)
	})
}