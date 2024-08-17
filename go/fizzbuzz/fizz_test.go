package fizz

import (
	"testing"
	"reflect"
)

func handleError(got, want []string, t *testing.T){
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFizzBuzz(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := FizzBuzz(3)
		want := []string{"1", "2", "Fizz"}
		handleError(got, want, t)
	})

	t.Run("", func(t *testing.T){
		got := FizzBuzz(5)
		want := []string{"1", "2", "Fizz", "4", "Buzz"}
		handleError(got, want, t)
	})

	t.Run("", func(t *testing.T){
		got := FizzB
	})
}