package validAnagram

import "testing"

func handleError(got, want bool, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestValidAnagram(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := ValidAnagram("car", "rac")
		want := true

		handleError(got, want, t)
	})

	t.Run("is not anagram", func(t *testing.T){
		got := ValidAnagram("rabbit", "tree")
		want := false

		handleError(got, want, t)
	})

	t.Run("is not anagram of equal length", func(t *testing.T){
		got := ValidAnagram("rabbit", "apples")
		want := false

		handleError(got, want, t)
	})


}