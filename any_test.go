package fstord

import "testing"

func TestAnyMap(t *testing.T) {
	mp := map[int]string{
		1: "a",
		3: "abc",
		5: "abcde",
	}

	result := Any(mp, func(a int, b string) bool {
		return len(b) != a
	})

	if result != false {
		t.Fatalf("expected %v should be false", result)
	}
}

func TestAnySlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	result := Any(slice, func(i int) bool {
		return i%2 == 1
	})

	if result != true {
		t.Fatalf("expected %v should be true", result)
	}
}
