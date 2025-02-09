package list

import (
	"cmp"
)

type node[T cmp.Ordered] struct {
	next  *node[T]
	value T
}

func newNode[T cmp.Ordered]() *node[T] {
	n := new(node[T])
	return n
}

type LinkedList[T cmp.Ordered] struct {
	head *node[T]
	tail *node[T]
	size int
}

// Constructs an empty LinkedList.
func NewLinkedList[T cmp.Ordered]() LinkedList[T] {
	ll := LinkedList[T]{
		head: newNode[T](),
		tail: nil,
		size: 0,
	}
	ll.tail = ll.head

	return ll
}

func (ll *LinkedList[T]) Add(element T) {
	ll.tail.value = element
	ll.tail.next = newNode[T]()
	ll.tail = ll.tail.next
	ll.size += 1
}

// Appends all of the elements in the specified collection to the end of this list,
// in the order that they are returned by the specified collection's iterator.
func (ll *LinkedList[T]) AddAll(elements ...T) {
	for i := 0; i < len(elements); i++ {
		ll.Add(elements[i])
	}
}

// Removes all of the elements from this list.
func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
}

// Returns the element at the specified position in this list.
func (ll *LinkedList[T]) Get(index int) T {
	n := ll.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n.value
}

// Returns an array containing all of the elements
// in this list in proper sequence (from first to last element).
func (ll *LinkedList[T]) ToArray() []T {
	arr := make([]T, 0)
	n := ll.head
	for i := 0; i < ll.size; i++ {
		arr = append(arr, n.value)
		n = n.next
	}
	return arr
}

// Removes the element at the specified position in this list.
func (ll *LinkedList[T]) Remove(index int) T {
	
	if ll.size == 0 {
		panic("list is empty.\n")
	}

	if index == 0 {
		val := ll.head.value
		ll.head = ll.head.next
		return val
	}

	n := ll.head
	var prev *node[T] = nil

	for i := 0; i < index; i++ {
		prev = n
		n = n.next
	}
	prev.next = nil
	ll.size -= 1
	return n.value
}


// Pops an element from the stack represented by this list.
func (ll *LinkedList[T]) Pop() T {
	return ll.Remove(ll.size-1)
}

// Returns the number of elements in this list.
func (ll LinkedList[T]) Size() int {
	return ll.size
}

