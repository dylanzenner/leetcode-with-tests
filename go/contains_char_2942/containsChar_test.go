package containschar

import (
	"testing"
	"reflect"
)

func handleError(got, want []int, t * testing.T){
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestContainsChar(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := ContainsChar([]string{"golang", "google"}, 'o')
		want := []int{0, 1}
		handleError(got, want, t)
	})

	t.Run("no output", func(t *testing.T){
		got := ContainsChar([]string{"abc", "def"}, 'z')
		want := []int{}
		handleError(got, want, t)
	})
}