package stableMtns

import (
	"testing"
	"reflect"
)

func TestStableMtns(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := StableMtns([]int{1, 2, 3, 4, 5}, 2)
		want := []int{3, 4}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := StableMtns([]int{10, 1, 10, 1, 10}, 3)
		want := []int{1, 3}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})
}