package fstord

import "testing"

func TestFilterMap(t *testing.T) {
	mp := map[int]string{
		1: "a",
		3: "ab",
		5: "abcde",
	}

	expected := map[int]string{
		1: "a",
		5: "abcde",
	}

	result := Filter(mp, func(a int, b string) bool {
		return len(b) == a
	}).(map[int]string)

	for k, v := range expected {
		if v != result[k] {
			t.Fatalf("expected %v to eq %v", expected, result)
		}
	}
}

func TestFilterSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	expected := []int{1, 3, 5}

	result := Filter(slice, func(i int) bool {
		return i%2 == 1
	}).([]int)

	for k, v := range expected {
		if v != result[k] {
			t.Fatalf("expected %v to eq %v", expected, result)
		}
	}
}

func TestFilterEmptySlice(t *testing.T) {
	slice := []int{}

	result := Filter(slice, func(i int) bool {
		return i%2 == 1
	}).([]int)

	if len(result) != 0 {
		t.Fatalf("expected %v to be empty", result)
	}
}
