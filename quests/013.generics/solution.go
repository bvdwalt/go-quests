package generics

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push adds a value to the end of the list.
func (lst *List[T]) Push(v T) {
	newElem := &element[T]{val: v}

	if lst.head == nil {
		lst.head = newElem
		lst.tail = newElem
		return
	}
	lst.tail.next = newElem
	lst.tail = newElem
}

// Pop removes the last element and returns it.
func (lst *List[T]) Pop() (T, bool) {
	var last = lst.tail

	if last == nil {
		var zero T
		return zero, false
	}

	if lst.head == lst.tail {
		lst.head = nil
		lst.tail = nil
		return last.val, true
	}

	current := lst.head
	for current.next != lst.tail {
		current = current.next
	}

	current.next = nil
	lst.tail = current

	return last.val, true
}

// AllElements returns a slice of all elements in the list.
func (lst *List[T]) AllElements() []T {
	result := []T{}

	for current := lst.head; current != nil; current = current.next {
		result = append(result, current.val)
	}

	return result
}
