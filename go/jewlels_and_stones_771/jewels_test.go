package jewels

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindJewels(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := FindJewels("aAzZ", "hwuelzauoAZ")
		want := 4

		handleError(got, want, t)
	})

	t.Run("no jewels", func(t *testing.T){
		got := FindJewels("", "kdjfkaj")
		want := 0
		handleError(got, want, t)
	})

	t.Run("all caps", func(t *testing.T){
		got := FindJewels("ZDH", "kjaehfZeiufDkwhH")
		want := 3
		handleError(got, want, t)
	})
}