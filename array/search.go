package array

import (
	"cmp"
)

// Search for the element in the array and gives the first index of the element
// otherwise -1
func LinearSearch[T comparable](elements []T, element T) int {

	for i := 0; i < len(elements); i++ {
		if elements[i] == element {
			return i
		}
	}

	return -1
}

// This is a reverse linear search algorithm
// Search for the element in the array and gives the first index of the element
// otherwise -1
func ReverseLinearSearch[T comparable](elements []T, element T) int {
	for i := len(elements) - 1; i >= 0; i-- {
		if elements[i] == element {
			return i
		}
	}
	return -1
}

// This is an experimental function
func ThreadSearch[T comparable](elements []T, element T) int {

	// Create a channel to receive the computed result.
	res := make(chan int)


	// Launch an anonymous function as a goroutine.
	go func(arr []T, ele T, ch chan int) {
		// Compute the square of n.
		result := LinearSearch(arr, ele)
		// Send the result back on the channel.
		ch <- result
	}(elements, element, res)


	go func(arr []T, ele T, ch chan int) {
		// Compute the square of n.
		result := ReverseLinearSearch(arr, ele)
		// Send the result back on the channel.
		ch <- result
	}(elements, element, res)

	// Wait for and receive the result from the channel.
	result, ok1 := <-res

	if result != -1 && ok1 {
		return result
	}


	result, ok2 := <-res

	if result != -1 && ok2 {
		return result
	}

	return -1
}

// This is a simple linear search which search from both side of an array
// Search for the element in the array and gives the first index of the element
// otherwise -1
func BiDirectionalSearch[T comparable](elements []T, element T) int {

	right := 0
	left := len(elements) - 1

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

	return -1
}

// Search for the element in the array and gives the first index of the element
// otherwise -1 .
//
// This is only aplicable if the array is sorted
func BinarySearch[T cmp.Ordered](elements []T, element T) int {
	low := 0
	high := len(elements)
	mid := int((high + low) / 2)

	for low != high {
		if elements[mid] < element {
			low = mid
		} else if elements[mid] > element {
			high = mid
		} else if element == elements[mid] {
			return mid
		}

		mid = int((high + low) / 2)
	}

	return -1
}
