package peaks

import (
	"testing"
	"reflect"
)

func TestPeaks(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := Peaks([]int{2, 4, 4})
		want := []int{}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := Peaks([]int{1, 4, 3, 8, 5})
		want := []int{1, 3}

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	})
}