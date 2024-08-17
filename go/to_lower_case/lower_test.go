package lower

import "testing"

func handleError(got, want string, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLower(t *testing.T){
	t.Run("base case", func(t *testing.T){
		got := Lower("HeLLo")
		want := "hello"
		handleError(got, want, t)
	})

	t.Run("all caps", func(t *testing.T){
		got := Lower("WORLD")
		want := "world"
		handleError(got, want, t)
	})

	t.Run("all lower", func(t *testing.T){
		got := Lower("awesome")
		want := "awesome"
		handleError(got, want, t)
	})
}