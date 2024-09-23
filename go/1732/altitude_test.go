package altitude

import "testing"

func TestAltitude(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := Altitude([]int{-5, 1, 5, 0, -7})
		want := 1

		if got != want {
			t.Errorf("got %v want % v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := Altitude([]int{-4, -3, -2, -1, 4, 3, 2})
		want := 0

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}