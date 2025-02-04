package tests

// import "testing"
import (
	"testing"
	"github.com/CodeOfSomnath/algox/list"
)

// testing the list add function working or not
func TestAddListFunction(t *testing.T)  {
	arr := list.NewArrayList[int]()
	arr.Add(10)
	arr.Add(20)

	a := []int{10, 20};
	b := arr.ToArray();
	for i := 0; i < arr.Size(); i++ {
		t.Logf("a[%d] = %d, b[%d]= %d", i, a[i], i, b[i]);
		if a[i] != b[i] {
			t.Errorf("%d != %d", a[i], b[i]);
		}
	}
}
