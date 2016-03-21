package fstord

import (
	"strconv"
	"testing"
)

func TestReduceSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	expected := "1491625"

	result := Reduce(slice, func(s string, x int) string {
		return s + strconv.Itoa(x*x)
	}, "").(string)

	if result != expected {
		t.Fatalf("expected %v to eq %v", expected, result)
	}
}

func TestReduceStruct(t *testing.T) {
	type ele struct {
		v int
		w int
	}
	slice := []ele{
		ele{1, 2},
		ele{3, 4},
	}

	expected := 14

	result := Reduce(slice, func(a int, e ele) int {
		return a + e.v*e.w
	}, 0).(int)

	if result != expected {
		t.Fatalf("expected %v to eq %v", expected, result)
	}
}

func TestReduceEmptySlice(t *testing.T) {
	slice := []int{}

	expected := ""

	result := Reduce(slice, func(s string, x int) string {
		return s + strconv.Itoa(x*x)
	}, "").(string)

	if result != expected {
		t.Fatalf("expected %v to eq %v", expected, result)
	}
}

func TestReduceMap(t *testing.T) {
	mp := map[int]string{
		1: "abc",
		2: "eft",
		3: "hij",
	}

	expected := 6

	result := Reduce(mp, func(s, k int, v string) int { return s + k }, 0).(int)

	if result != expected {
		t.Fatalf("expected %v to eq %v", expected, result)
	}
}
