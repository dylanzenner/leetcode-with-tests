package subseq

import "testing"

func handleError(got, want bool, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSubSeq(t *testing.T){
	t.Run("is subseq", func(t *testing.T){
		got := SubSeq("car", "abcardef")
		want := true
		handleError(got, want, t)
	})

	t.Run("is not subseq", func(t *testing.T){
		got := SubSeq("abc", "def")
		want := false
		handleError(got, want, t)
	})
}