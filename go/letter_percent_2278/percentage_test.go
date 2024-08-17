package percentage 

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPercentage(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Percentage("letter", 't')
		want := 33
		handleError(got, want, t)
	})

	t.Run("zero", func(t *testing.T){
		got := Percentage("zzzzz", 'k')
		want := 0
		handleError(got, want, t)
	})
}