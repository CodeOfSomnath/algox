package tests

import (
	"testing"

	"github.com/CodeOfSomnath/algox/array"
)

func TestArrayRemove(t *testing.T) {
	arr := []int{2, 3, 4, 10, 15}
	arr = array.Remove(arr, 10)
	if arr[3] != 15 {
		t.Errorf("array.Remove() function is not working\n")
	}
}

func TestArrayRemoveAt(t *testing.T) {
	arr := []int{2, 3, 4, 10, 15}
	arr = array.RemoveAt(arr, 3)
	if arr[3] != 15 {
		t.Errorf("array.Remove() function is not working\n")
	}
}

func TestArrayRemoveAll(t *testing.T) {
	arr := []int{2, 3, 4, 10, 10, 15}
	arr = array.RemoveAll(arr, 10)
	if arr[3] != 15 {
		t.Errorf("array.Remove() function is not working\n")
		t.Errorf("array= %v\n", arr)
	}
}
