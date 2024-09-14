package runningsum

import (
	"testing"
	"reflect"
)

func handleError(got, want []int, t *testing.T){
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRunningSum(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := RunningSum([]int{1, 2, 3, 4})
		want := []int{1, 3, 6, 10}
		handleError(got, want, t)
	})

	t.Run("second test", func(t *testing.T){
		got := RunningSum([]int{1, 1, 1, 1, 1})
		want := []int{1, 2, 3, 4, 5}
		handleError(got, want, t)
	})
}