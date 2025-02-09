package array


// Search for the element in the arraylist and gives the first index of the element
// otherwise -1
func LinearSearch[T comparable](elements []T,element T) int {
	
	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return i
		}
	}

	return -1;
}

// This is a exprimental searching algorithm.
func LinearSearchX[T comparable](elements []T,element T) int {
	
	right := 0
	left := len(elements)-1

	for right <= left {
		
		if elements[right] == element {
			return right
		}

		if elements[left] == element {
			return left
		}

		right += 1
		left -= 1
	}

	return -1;
}