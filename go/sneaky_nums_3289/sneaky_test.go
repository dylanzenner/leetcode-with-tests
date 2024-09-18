package sneakyNums

import (
	"testing"
	"reflect"
)

func TestSneakyNums(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := SneakyNums([]int{0, 1, 1, 0})
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := SneakyNums([]int{0, 3, 2, 1, 3, 2})
		want := []int{2, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}