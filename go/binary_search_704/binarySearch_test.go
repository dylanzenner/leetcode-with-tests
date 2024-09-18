package binarySearch

import "testing"

func TestBinarySearch(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := BinarySearch([]int{1, 2, 3, 4, 5, 6}, 3)
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := BinarySearch([]int{0, -1, 4, 10, 100, -300}, 200)
		want := -1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 3", func(t *testing.T){
		got := BinarySearch([]int{-100, -150, -200, -300, 100, 200, 300}, -150)
		want := 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}