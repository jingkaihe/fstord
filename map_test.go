package fstord

import (
	"strconv"
	"testing"
)

func TestMapMap(t *testing.T) {
	mp := map[int]int{
		1: 2,
		3: 4,
		5: 6,
	}

	expected := map[int]string{
		1: "3",
		3: "5",
		5: "7",
	}

	result := Map(mp, func(a, b int) string {
		return strconv.Itoa(b + 1)
	}).(map[int]string)

	for k, v := range expected {
		if v != result[k] {
			t.Fatalf("expected %v to eq %v", expected, result)
		}
	}
}

func TestMapSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	expected := []string{"1", "4", "9", "16", "25"}

	result := Map(slice, func(x int) string {
		return strconv.Itoa(x * x)
	}).([]string)

	for k, v := range expected {
		if v != result[k] {
			t.Fatalf("expected %v to eq %v", expected, result)
		}
	}
}
