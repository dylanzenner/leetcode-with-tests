package addTwoInts

import "testing"

func TestAddTwoInts(t *testing.T){

	t.Run("add two positive ints", func(t *testing.T){
		got := AddTwoInts(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("add a negative and a positive int", func(t *testing.T){
		got := AddTwoInts(2, -1)
		want := 1

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("add two negative ints", func(t *testing.T){
		got := AddTwoInts(-2, -2)
		want := -4

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("add a positive int and zero", func(t *testing.T){
		got := AddTwoInts(5, 0)
		want := 5

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("add a negative int and zero", func(t *testing.T){
		got := AddTwoInts(-5, 0)
		want := -5

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("add two zeros", func(t *testing.T){
		got := AddTwoInts(0, 0)
		want := 0

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}