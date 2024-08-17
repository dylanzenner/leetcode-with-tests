package convert

import (
	"testing"
	"reflect"
)

func handleError(got, want []float64, t *testing.T){
	if !reflect.DeepEqual(got, want){
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestConvert(t *testing.T){
	t.Run("base test", func(t *testing.T){
		got := Convert(36.50)
		want := []float64{309.65000, 97.70000}
		handleError(got, want, t)
	})

	t.Run("test 2", func(t *testing.T){
		got := Convert(122.11)
		want := []float64{395.26000, 251.79800}
		handleError(got, want, t)
	})
}