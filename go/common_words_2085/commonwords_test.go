package commonwords

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCommonWords(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := CommonWords([]string{"golang", "is", "awesome"}, []string{"golang", "is", "cool"})
		want := 2
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := CommonWords([]string{"aaa", "a", "b"}, []string{"aaa", "bbb", "zzz"})
		want := 1
		handleError(got, want, t)
	})
}