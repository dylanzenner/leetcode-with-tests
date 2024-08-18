package equaloccurrences

import "testing"

func handleError(got, want bool, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestEqualOccurrences(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := EqualOccurrences("abacbc")
		want := true
		handleError(got, want, t)
	})

	t.Run("fail case", func(t *testing.T){
		got := EqualOccurrences("abbccabbccc")
		want := false
		handleError(got, want, t)
	})
}