package isacronym

import "testing"

func handleError(got, want bool, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsAcronym(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := IsAcronym([]string{"golang", "is", "cool"}, "gic")
		want := true
		handleError(got, want, t)
	})

	t.Run("false test", func(t *testing.T){
		got := IsAcronym([]string{"this", "is", "incorrect"}, "abc")
		want := false
		handleError(got, want, t)
	})
}