package tests

import (
	"testing"

	"github.com/CodeOfSomnath/algox/list"
)


func TestLinkedListBasic(t *testing.T)  {
	linkedList := list.NewLinkedList[int]()
	linkedList.Add(10)
	val := linkedList.Get(0)
	if val != 10 {
		t.Errorf("LinkedList.Get() function is not working, %v != %v\n", val, 10)
	}

	linkedList.AddAll(2,3,4)
	arr := linkedList.ToArray()
	b := []int{10, 2,3,4}
	for i := 0; i < len(arr); i++ {
		if arr[i] != b[i] {
			t.Errorf("LinkedList.ToArray() is not working, index=%v, %v!=%v\n", i, arr[i], b[i])
		}
		t.Logf("arr[%v] = %v, b[%v] = %v\n", i, arr[i], i, b[i])
	}
}