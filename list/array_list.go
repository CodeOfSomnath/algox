package list

type ArrayList[T comparable] struct {
	
	array []T
}


func NewArrayList[T comparable]() ArrayList[T] {
	return ArrayList[T] {
		array: make([]T, 0),
	}
}

// Add a element to the ArrayList
func (arr * ArrayList[T]) Add(value T)  {
	arr.array = append(arr.array, value)
}

// Inserts the specified element at the specified position in this list.
func (arr * ArrayList[T]) AddAt(index int, value T)  {
	
	currentElement := value
	// currentIndex := index
	var temp T

	for i := index; i < len(arr.array); i++ {
		temp = arr.array[i]
		arr.array[i] = currentElement
		currentElement = temp
	}

	arr.array = append(arr.array, currentElement)

}

// Appends all of the elements in the specified collection to the end of this list,
// in the order that they are returned by the specified collection's Iterator.
func (arr *  ArrayList[T]) AddAll(values ...T) bool {
	isSuccessFull := true

	for i := 0; i < len(values); i++ {
		arr.Add(values[i])
	}

	return isSuccessFull;
}

// Inserts all of the elements in the specified collection into this list,
// starting at the specified position.
func (arr *  ArrayList[T]) AddAllAt(index int, values ...T) bool {
	isSuccessFull := true

	for i := 0; i < len(values); i++ {
		arr.AddAt(index, values[i])
		index += 1
	}

	return isSuccessFull;
}

// Removes all of the elements from this list.
func (arr *  ArrayList[T]) Clear()  {
	arr.array = make([]T, 0)
}

// Returns a shallow copy of this ArrayList instance.
func (arr *  ArrayList[T]) Clone() ArrayList[T] {
	copyArrayList := *arr
	return copyArrayList;
}

// Returns true if this list contains the specified element.
func (arr *  ArrayList[T]) Contains(value T) bool  {
	return arr.IndexOf(value) != -1;
}

// Returns the index of the first occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (arr *  ArrayList[T]) IndexOf(value T) int  {
	for i := 0; i < len(arr.array); i++ {
		if arr.array[i] == value {
			return i;
		}
	}
	return -1;
}

// Returns the index of the last occurrence of the specified element in this list,
// or -1 if this list does not contain the element.
func (arr * ArrayList[T]) LastIndexOf(value T) int {
	for i := len(arr.array)-1; i >= 0; i-- {
		if arr.array[i] == value {
			return i;
		}
	}
	return -1;
}


// Return a element from the ArrayList
func (arr * ArrayList[T]) Get(index int) T {
	return arr.array[index];
}

// Returns true if this list contains no elements.
func (arr * ArrayList[T]) IsEmpty() bool {
	return len(arr.array) == 0
}

// Removes the element at the specified position in this list.
func (arr * ArrayList[T]) RemoveAt(index int) T {
	value := arr.array[index]
	arr.array = append(arr.array[:index], arr.array[index+1:]...)
	return value
}

// Removes the first occurrence of the specified element from this list, if it is present.
func (arr * ArrayList[T]) Remove(value T) bool {
	index := arr.IndexOf(value)
	if index == -1 {
		return false
	}

	// remove
	arr.RemoveAt(index)

	return true
}

// Removes from this list all of its elements that are contained in the specified collection.
func (arr * ArrayList[T]) RemoveAll(values ...T) bool {
	isChanged := false

	for i := 0; i < len(values); i++ {
		val := arr.Remove(values[i])
		isChanged = isChanged || val
	}
	return isChanged;
}

// Replaces the element at the specified position in this list with the specified element.
func (arr * ArrayList[T]) Set(index int, value T) T {
	val := arr.array[index];
	arr.array[index] = value
	return val
}

// Returns the number of elements in this list.
func (arr ArrayList[T]) Size() int {
	return len(arr.array)
}


// Returns an array containing all of the elements
// in this list in proper sequence (from first to last element).
func (arr ArrayList[T]) ToArray() []T {
	return arr.array
}

