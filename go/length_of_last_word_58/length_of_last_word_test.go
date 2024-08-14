package lengthOfLastWord

import "testing"

func TestLengthOfLastWord(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := LengthOfLastWord("   Hello World   ")
		want := 5

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("test no spaces at beginning", func(t *testing.T){
		got := LengthOfLastWord("This is another string   ")
		want := 6

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("no spaces at end of string", func(t *testing.T){
		got := LengthOfLastWord("This   is a string  with no spaces at the end")
		want := 3

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("empty string test", func(t *testing.T){
		got := LengthOfLastWord("")
		want := 0

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}