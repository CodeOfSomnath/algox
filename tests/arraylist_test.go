package tests

// import "testing"
import (
	"github.com/CodeOfSomnath/algox/list"
	"testing"
)

// testing the list add function working or not
func TestAddListFunction(t *testing.T) {
	arr := list.NewArrayList[int]()
	arr.Add(10)
	arr.Add(20)

	a := []int{10, 20}
	b := arr.ToArray()
	for i := 0; i < arr.Size(); i++ {
		t.Logf("a[%d] = %d, b[%d]= %d", i, a[i], i, b[i])
		if a[i] != b[i] {
			t.Errorf("%d != %d", a[i], b[i])
		}
	}
}

func ArrayCheker[T comparable](a []T, b []T) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestAddFunctions(t *testing.T) {
	arr := list.NewArrayList[int]()
	a := []int{10, 34, 23}

	arr.AddAll(a...)

	//checking if all elements are add
	if !ArrayCheker(arr.ToArray(), a) {
		t.Errorf("%v != %v\n", arr.ToArray(), a)
	}
	// Removing all added elements
	arr.RemoveAll(a...)

	if len(arr.ToArray()) != 0 {
		t.Errorf("ArrayList.RemoveAll() is not working, array = %v\n", arr.ToArray())
	}

}

func TestIndexFunctions(t *testing.T) {
	arr := list.NewArrayList[int]()
	arr.AddAll(10, 20, 30)
	if arr.IndexOf(20) != 1 {
		t.Errorf("index = %v\n, is not 1", arr.IndexOf(20))
	}
}

func TestSort(t *testing.T) {
	arr := list.NewArrayList[int]()
	arr.AddAll(7, 4, 10)
	arr.Sort()
	if !ArrayCheker(arr.ToArray(), []int{4, 7, 10}) {
		t.Errorf("ArrayList.Sort() not working, %v\n", arr.ToArray())
	}
}


func TestArrayListStringMethod(t *testing.T)  {
	arr := list.NewArrayList[int]()
	arr.AddAll(2,3,4,5)
	t.Logf("%v\n", arr)
}
