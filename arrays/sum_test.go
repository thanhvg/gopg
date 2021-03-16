package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		wanted := 15
		if got != wanted {
			t.Errorf("Expect %d got %d", wanted, got)
		}
	})

	t.Run("Collection of 3 number", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		wanted := 6
		if got != wanted {
			t.Errorf("Expect %d got %d", wanted, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2})
	wanted := []int{6, 3}
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("Expect %v got %v", wanted, got)
	}
}

func TestSumTail(t *testing.T) {
	t.Run("Ok with proper slices", func(t *testing.T) {
		got := SumTail([]int{1, 2, 3}, []int{1, 2})
		wanted := []int{5, 2}
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("Expect %v got %v", wanted, got)
		}
	})
	t.Run("Ok with one item slices", func(t *testing.T) {
		got := SumTail([]int{3}, []int{2})
		wanted := []int{0, 0}
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("Expect %v got %v", wanted, got)
		}
	})

	t.Run("Ok with empty slices", func(t *testing.T) {
		got := SumTail([]int{}, []int{})
		wanted := []int{0, 0}
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("Expect %v got %v", wanted, got)
		}
	})
}
