package its_test

import (
	"aoc-lib/its"
	"iter"
	"slices"
	"testing"
)

func TestWindow2(t *testing.T) {
	t.Run("window empty", func(t *testing.T) {
		data := []int{}
		next, stop := iter.Pull2(its.Window2(slices.Values(data)))
		defer stop()

		_, _, ok := next()
		if ok {
			t.Fatal("There should be no value")
		}
	})

	t.Run("window 1 value", func(t *testing.T) {
		data := []int{1}
		next, stop := iter.Pull2(its.Window2(slices.Values(data)))
		defer stop()

		_, _, ok := next()
		if ok {
			t.Fatal("There should be no value")
		}
	})

	t.Run("window 2 values", func(t *testing.T) {
		data := []int{1, 2}
		next, stop := iter.Pull2(its.Window2(slices.Values(data)))
		defer stop()
		expectedFirst, expectedSecond := 1, 2

		first, second, ok := next()
		if !ok {
			t.Fatal("There should be a value")
		}

		if expectedFirst != first {
			t.Fatalf("Error: Expected First value to be %d, actual %d", expectedFirst, first)
		}

		if expectedSecond != second {
			t.Fatalf("Error: Expected Second value to be %d, actual %d", expectedSecond, second)
		}

		_, _, ok = next()
		if ok {
			t.Fatal("There should be no value after first pull")
		}
	})

	t.Run("window 3 values", func(t *testing.T) {
		data := []int{1, 2, 3}
		next, stop := iter.Pull2(its.Window2(slices.Values(data)))
		defer stop()

		expectedFirst, expectedSecond := 1, 2

		first, second, ok := next()
		if !ok {
			t.Fatal("There should be a value")
		}

		if expectedFirst != first {
			t.Fatalf("Error: Expected First value to be %d, actual %d", expectedFirst, first)
		}

		if expectedSecond != second {
			t.Fatalf("Error: Expected Second value to be %d, actual %d", expectedSecond, second)
		}

		expectedFirst, expectedSecond = 2, 3

		first, second, ok = next()
		if !ok {
			t.Fatal("There should be a value after the first")
		}

		if expectedFirst != first {
			t.Fatalf("Error: Expected First value of second pull to be %d, actual %d", expectedFirst, first)
		}

		if expectedSecond != second {
			t.Fatalf("Error: Expected Second value  of second pull to be %d, actual %d", expectedSecond, second)
		}

		_, _, ok = next()
		if ok {
			t.Fatal("There should be no value after second pull")
		}
	})
}

func TestAllCombinations(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		slice := []int{}
		expected := []its.Combination[int]{}

		actual := slices.Collect(its.AllCombinations(slice, false))

		if len(expected) != len(actual) {
			t.Fatalf("Error: expected Length to be %d, actual %d", len(expected), len(actual))
		}

		if !its.All2(its.ZipSlices(expected, actual), func(expected, actual its.Combination[int]) bool {
			return expected == actual
		}) {
			t.Fatalf("Error: expected %+v, actual %+v", expected, actual)
		}
	})

	t.Run("one value not self", func(t *testing.T) {
		slice := []int{1}
		expected := []its.Combination[int]{}

		actual := slices.Collect(its.AllCombinations(slice, false))

		if len(expected) != len(actual) {
			t.Fatalf("Error: expected Length to be %d, actual %d", len(expected), len(actual))
		}

		if !its.All2(its.ZipSlices(expected, actual), func(expected, actual its.Combination[int]) bool {
			return expected == actual
		}) {
			t.Fatalf("Error: expected %+v, actual %+v", expected, actual)
		}
	})

	t.Run("one value self", func(t *testing.T) {
		slice := []int{1}
		expected := []its.Combination[int]{{1, 1}}

		actual := slices.Collect(its.AllCombinations(slice, true))

		if len(expected) != len(actual) {
			t.Fatalf("Error: expected Length to be %d, actual %d", len(expected), len(actual))
		}

		if !its.All2(its.ZipSlices(expected, actual), func(expected, actual its.Combination[int]) bool {
			return expected == actual
		}) {
			t.Fatalf("Error: expected %+v, actual %+v", expected, actual)
		}
	})

	t.Run("two values not self", func(t *testing.T) {
		slice := []int{1, 2}
		expected := []its.Combination[int]{{1, 2}}

		actual := slices.Collect(its.AllCombinations(slice, false))

		if len(expected) != len(actual) {
			t.Fatalf("Error: expected Length to be %d, actual %d", len(expected), len(actual))
		}

		if !its.All2(its.ZipSlices(expected, actual), func(expected, actual its.Combination[int]) bool {
			return expected == actual
		}) {
			t.Fatalf("Error: expected %+v, actual %+v", expected, actual)
		}
	})

	t.Run("two values self", func(t *testing.T) {
		slice := []int{1, 2}
		expected := []its.Combination[int]{{1, 1}, {1, 2}, {2, 2}}

		actual := slices.Collect(its.AllCombinations(slice, true))

		if len(expected) != len(actual) {
			t.Fatalf("Error: expected Length to be %d, actual %d", len(expected), len(actual))
		}

		if !its.All2(its.ZipSlices(expected, actual), func(expected, actual its.Combination[int]) bool {
			return expected == actual
		}) {
			t.Fatalf("Error: expected %+v, actual %+v", expected, actual)
		}
	})

	t.Run("five values not self", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := []its.Combination[int]{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5}}

		actual := slices.Collect(its.AllCombinations(slice, false))

		if len(expected) != len(actual) {
			t.Fatalf("Error: expected Length to be %d, actual %d", len(expected), len(actual))
		}

		if !its.All2(its.ZipSlices(expected, actual), func(expected, actual its.Combination[int]) bool {
			return expected == actual
		}) {
			t.Fatalf("Error: expected %+v, actual %+v", expected, actual)
		}
	})

	t.Run("five values self", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := []its.Combination[int]{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 2}, {2, 3}, {2, 4}, {2, 5}, {3, 3}, {3, 4}, {3, 5}, {4, 4}, {4, 5}, {5, 5}}

		actual := slices.Collect(its.AllCombinations(slice, true))

		if len(expected) != len(actual) {
			t.Fatalf("Error: expected Length to be %d, actual %d", len(expected), len(actual))
		}

		if !its.All2(its.ZipSlices(expected, actual), func(expected, actual its.Combination[int]) bool {
			return expected == actual
		}) {
			t.Fatalf("Error: expected %+v, actual %+v", expected, actual)
		}
	})
}
