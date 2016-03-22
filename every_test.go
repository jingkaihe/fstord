package fstord

import "testing"

func TestEveryMap(t *testing.T) {
	mp := map[int]string{
		1: "a",
		3: "abc",
		5: "abcde",
	}

	result := Every(mp, func(a int, b string) bool {
		return len(b) == a
	})

	if result != true {
		t.Fatalf("expected %v should be true", result)
	}
}

func TestEverySlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	result := Every(slice, func(i int) bool {
		return i%2 == 1
	})

	if result != false {
		t.Fatalf("expected %v should be false", result)
	}
}
