package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of any", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d, got %d, passing %v", want, got, numbers)
		}
	})

}
