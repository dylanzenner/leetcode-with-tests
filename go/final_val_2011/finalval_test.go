package finalval

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFinalVal(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := FinalVal([]string{"--X", "X++", "X++"})
		want := 1
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := FinalVal([]string{"++X", "++X", "X++"})
		want := 3
		handleError(got, want, t)
	})
}