package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, got, expected, numbers)
	})

}

func assertCorrectMessage(t testing.TB, got, want int, value []int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d, got %d, value passed in %v", got, want, value)
	}
}
