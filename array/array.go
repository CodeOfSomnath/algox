// Package array gives all array related functions like search, copy, clone, remove, etc.
package array

import "fmt"

// This function removes the element from the array in go
func Remove[T comparable](elements []T, element T) []T {
	i := LinearSearch(elements, element)

	if i == -1 {
		return elements
	}

	elements = append(elements[:i], elements[i+1:]...)
	return elements
}

// This function removes the element at that index from the array in go
func RemoveAt[T comparable](elements []T, i int) []T {
	if i == -1 {
		return elements
	}

	elements = append(elements[:i], elements[i+1:]...)
	return elements
}

// This function removes all the elements from the array in go
func RemoveAll[T comparable](elements []T, element T) []T {
	i := LinearSearch(elements, element)
	fmt.Printf("i=%v\n", i)
	for i != -1 {
		elements = append(elements[:i], elements[i+1:]...)
		i = LinearSearch(elements, element)
	}

	return elements
}
