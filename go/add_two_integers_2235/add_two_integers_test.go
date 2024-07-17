package addtwointegers2235

import "testing"

func TestAddTwoIntegers(t *testing.T){
	t.Run("Test two positive integers", func(t *testing.T) {
		got := AddTwoIntegers(10, 10)
		want := 20

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Test two negative numbers", func(t *testing.T){
		got := AddTwoIntegers(-5, -10)
		want := -15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Test a positive and a negative integer", func(t *testing.T){
		got := AddTwoIntegers(10, -5)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Test zero and a positive number", func(t *testing.T){
		got := AddTwoIntegers(0, 10)
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Test zero and a negative number", func(t *testing.T){
		got := AddTwoIntegers(0, -5)
		want := -5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}