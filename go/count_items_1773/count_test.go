package count

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCount(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Count([][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}, "color", "silver")
		want := 1
		handleError(got, want, t)
	})
}