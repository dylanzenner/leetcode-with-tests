package mostwords

import "testing"

func handleError(got, want int, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMostWords(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := MostWords([]string{"golang is cool", "leetcode is interesting", "python is alright"})
		want := 3
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := MostWords([]string{"this is another sentence", "did anyone else go to reinvent last year"})
		want := 8
		handleError(got, want, t)
	})
}