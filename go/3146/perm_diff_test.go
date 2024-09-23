package permDiff

import "testing"

func TestPermDiff(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := PermDiff("abc", "bac")
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := PermDiff("abcde", "edbac")
		want := 12

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}