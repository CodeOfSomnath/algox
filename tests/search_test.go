package tests

import (
	"math/rand"
	"testing"

	"github.com/CodeOfSomnath/algox/array"
)




func TestBiDirectionalSearch(t *testing.T) {
	size := 20000
	x := size / 2
	arr := make([]int, size) // Create an array with 1000 elements
    for i := range arr {
        arr[i] = rand.Intn(size+1)
    }
	value := int(size/3 * 2)
	arr[x] = value
	index := array.BiDirectionalSearch(arr, value)
	t.Logf("value in %v =  %v", x, arr[x])
	if index == -1 {
		t.Logf("value = %v not found\n", value)
		return
	} else {
		t.Logf("value = %v found\n", arr[index])
	}

	if arr[index] != value {
		t.Errorf("arr[%v] = %v != %v \n", index, arr[index], value)
	}
	
}


func TestLinearSearch(t *testing.T) {
	size := 20000
	x := size / 2
	arr := make([]int, size) // Create an array with 1000 elements
    for i := range arr {
        arr[i] = rand.Intn(size+1)
    }
	value := int(size/3 * 2)
	arr[x] = value
	index := array.LinearSearch(arr, value)
	t.Logf("value in %v =  %v", x, arr[x])
	if index == -1 {
		t.Logf("value = %v not found\n", value)
		return
	} else {
		t.Logf("value = %v found\n", arr[index])
	}

	if arr[index] != value {
		t.Errorf("arr[%v] = %v != %v \n", index, arr[index], value)
	}
	
}




func TestBinarySearch(t * testing.T)  {
	arr := []int {2,3,4,10,15,25}
	i := len(arr)-1
	index := array.BinarySearch(arr, arr[i])
	if index == -1 {
		t.Logf("element not found\n")
	} else if index != i {
		t.Errorf("array.BinarySearch is not working\n")
	}   else {
		t.Logf("%v found at %v\n", arr[index], index)
	}
}


func TestReverseLinearSearch(t *testing.T)  {
	arr := []int {2,3,4,10,15,25}
	i := len(arr)-1
	index := array.ReverseLinearSearch(arr, arr[i])
	if index == -1 {
		t.Logf("element not found\n")
	} else if index != i {
		t.Errorf("array.BinarySearch is not working\n")
	}   else {
		t.Logf("%v found at %v\n", arr[index], index)
	}
}


func TestThreadSearch(t *testing.T)  {
	arr := []int {2,3,4,10,15,25}
	i := 0
	index := array.ThreadSearch(arr, arr[i])
	if index == -1 {
		t.Logf("element not found\n")
	} else if index != i {
		t.Errorf("array.BinarySearch is not working\n")
	}   else {
		t.Logf("%v found at %v\n", arr[index], index)
	}
}


func TestThreadSearchSecond(t *testing.T)  {
	size := 20000
	x := size / 2
	arr := make([]int, size) // Create an array with 1000 elements
    for i := range arr {
        arr[i] = rand.Intn(size+1)
    }
	value := int(size/3 * 2)
	arr[x] = value
	index := array.ThreadSearch(arr, value)
	t.Logf("value in %v =  %v", x, arr[x])
	if index == -1 {
		t.Logf("value = %v not found\n", value)
		return
	} else {
		t.Logf("value = %v found\n", arr[index])
	}

	if arr[index] != value {
		t.Errorf("arr[%v] = %v != %v \n", index, arr[index], value)
	}
	
}