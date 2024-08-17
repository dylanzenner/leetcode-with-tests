package equal

import "testing"

func handleError(got, want bool, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestEqual(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Equal([]string{"ab", "c", "def"}, []string{"abc", "def"})
		want := true

		handleError(got, want, t)
	})

	t.Run("not equal", func(t *testing.T){
		got := Equal([]string{"abc", "def"}, []string{"gh", "zy"})
		want := false
		handleError(got, want, t)
	})
}