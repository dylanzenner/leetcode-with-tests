package containsDuplicate

import (
	"testing"
	"fmt"
)

func handleError(got, want bool) string{
	if got != want {
		return fmt.Sprintf("got %v, want %v", got, want)
	}
	return ""
}

func TestContainsDuplicate(t *testing.T){
	t.Run("test no dups", func(t *testing.T){
		got := ContainsDuplicate([]int{1, 2, 3, 4})
		want := false

		handleError(got, want)
	})

	t.Run("test dups", func(t *testing.T){
		got := ContainsDuplicate([]int{1, 1, 3, 3, 4, 2, 7, 4})
		want := true

		handleError(got, want)
	})
}