package concat 

import (
	"testing"
	"reflect"
)

func handleError(got, want []int, t *testing.T){
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestConcat(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Concat([]int{1, 2, 1})
		want := []int{1, 2, 1, 1, 2, 1}
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := Concat([]int{1, 3, 2, 1})
		want := []int{1, 3, 2, 1, 1, 3, 2, 1}
		handleError(got, want, t)
	})
}