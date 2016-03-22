package fstord

import "testing"

func TestCountMap(t *testing.T) {
	mp := map[int]string{
		1: "a",
		3: "ab",
		5: "abcde",
	}

	result := Count(mp, func(a int, b string) bool {
		return len(b) == a
	})

	if result != 2 {
		t.Fatalf("expected %v should be 2", result)
	}
}

func TestCountSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	result := Count(slice, func(i int) bool {
		return i%2 == 1
	})

	if result != 3 {
		t.Fatalf("expected %v should be 3", result)
	}
}
