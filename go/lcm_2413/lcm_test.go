package lcm

import "testing"

func TestLcm(t *testing.T){
	t.Run("test 1", func(t *testing.T){
		got := Lcm(6)
		want := 6

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test 2", func(t *testing.T){
		got := Lcm(3)
		want := 6

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}