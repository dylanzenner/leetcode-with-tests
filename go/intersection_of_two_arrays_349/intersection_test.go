package intersect

import (
	"testing"
	"reflect"
)

func handleError(got, want []int, t *testing.T){
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInteresect(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Intersect([]int{1, 2, 2, 3}, []int{2, 2})
		want := []int{2}
		handleError(got, want, t)
	})

	t.Run("no intersection", func(t *testing.T){
		got := Intersect([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{}
		handleError(got, want, t)
	})
}