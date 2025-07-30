package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want, numbers)
	})
}

func assertCorrectMessage[T any](t testing.TB, got, want int, arr T) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d, given %v", got, want, arr)
	}
}
