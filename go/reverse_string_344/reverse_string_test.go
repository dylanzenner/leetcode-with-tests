package reverseString

import (
	"testing"
	"reflect"
)

func handleError(got, want []byte, t *testing.T){
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseString(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := ReverseString([]byte{'h', 'e', 'l', 'l', 'o'})
		want := []byte{'o', 'l', 'l', 'e', 'h'}

		handleError(got, want, t)
	})

	t.Run("empty input", func(t *testing.T){
		got := ReverseString([]byte{})
		want := []byte{}

		handleError(got, want, t)
	})

	t.Run("single space", func(t *testing.T){
		got := ReverseString([]byte{' '})
		want := []byte{' '}

		handleError(got, want, t)
	})
}