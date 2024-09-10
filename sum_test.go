package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, got, expected, numbers)
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Sum of many slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		assertDeepEqual(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Testing sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertDeepEqual(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 5}

		assertDeepEqual(t, got, want)
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2}, []int{0, 9})
	}
}

func assertCorrectMessage(t testing.TB, got, want int, value []int) {
	t.Helper()
	if got != want {
		t.Errorf("expected %d, got %d, value passed in %v", got, want, value)
	}
}

func assertDeepEqual(t testing.TB, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
