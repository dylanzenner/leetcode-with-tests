package defangIp

import "testing"

func handleError(got, want string, t *testing.T){
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDefangIp(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := DefangIp("1.1.1.1")
		want := "1[.]1[.]1[.]1"

		handleError(got, want, t)
	})
}