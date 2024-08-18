package panagram

import "testing"

func handleError(got, want bool, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPanagram(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Panagram("thequickbrownfoxjumpsoverthelazydog")
		want := true
		handleError(got, want, t)
	})

	t.Run("invalid test", func(t *testing.T){
		got := Panagram("golang")
		want := false
		handleError(got, want, t)
	})

	t.Run("invalid with 26 letters", func(t *testing.T){
		got := Panagram("thequickbrowkfoxjumpedoverthelazycat")
		want := false
		handleError(got, want, t)
	})
}