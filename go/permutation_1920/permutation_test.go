package permutation

import (
	"testing"
	"reflect"
)

func handleError(got, want []int, t *testing.T){
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPermutate(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Permutate([]int{0, 2, 1, 5, 3, 4})
		want := []int{0, 1, 2, 4, 5, 3}
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := Permutate([]int{5, 0, 1, 2, 3, 4})
		want := []int{4, 5, 0, 1, 2, 3}
		handleError(got, want, t)
	})
}