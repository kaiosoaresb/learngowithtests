package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll2([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkSumAll1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll1([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15})
	}
}

func BenchmarkSumAll2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll2([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{11, 12, 13, 14, 15})
	}
}

func TestSumAllTails(t *testing.T) {

	assertSum := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertSum(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertSum(t, got, want)
	})
}
