package uncommon

import (
	"testing"
	"reflect"
)

func handleError(got, want []string, t *testing.T){
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestUncommon(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Uncommon("hello world", "the world is round")
		want := []string{"hello", "the", "is", "round"}
		handleError(got, want, t)
	})

	t.Run("no uncommon words", func(t *testing.T){
		got := Uncommon("hello world", "hello world")
		want := []string{}

		handleError(got, want, t)
	})
}