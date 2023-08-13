package arrays

import (
	"reflect"
	"testing"
)

func Check(got int, want int, numbers []int, t *testing.T) {
	if got != want {
		t.Errorf("got: %d \n given: %d \n -numbers: %v", got, want, numbers)
	}

}

func TestSum(t *testing.T) {

	t.Run("collections of n numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		Check(got, want, numbers, t)

	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{0, 9})
	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
