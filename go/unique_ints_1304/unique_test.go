package unique

import "testing"

func handleError(got []int, want int, t *testing.T){
	count := 0
	for _, i := range got {
		count += i
	}

	if count != want {
		t.Errorf("got %v, want %v", got , want)
	}
}

func TestUnique(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Unique(5)
		want := 0
		handleError(got, want, t)
	})

	t.Run("large number", func(t *testing.T){
		got := Unique(100)
		want := 0
		handleError(got, want, t)
	})

	t.Run("larger num", func(t *testing.T){
		got := Unique(100000)
		want := 0
		handleError(got, want, t)
	})
}
