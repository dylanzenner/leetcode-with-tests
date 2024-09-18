package shuffle

import (
	"testing"
	"reflect"
)

func TestShuffle(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := Shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
		want := []int{2, 3, 5, 4, 1, 7}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	
	})

	t.Run("test 2", func(t *testing.T){
		got := Shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4)
		want := []int{1, 4, 2, 3, 3, 2, 4, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got , want)
		}
	})
}