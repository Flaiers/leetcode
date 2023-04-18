package stack

type (
	stack[T any] struct {
		top    *node[T]
		length int
	}
	node[T any] struct {
		value *T
		prev  *node[T]
	}
)

// New create a new stack
func New[T any]() *stack[T] {
	return &stack[T]{nil, 0}
}

// Len return the number of items in the stack
func (s *stack[T]) Len() int {
	return s.length
}

// Peek view the top item on the stack
func (s *stack[T]) Peek() *T {
	if s.length == 0 {
		return nil
	}

	return s.top.value
}

// Pop the top item of the stack and return it
func (s *stack[T]) Pop() *T {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *stack[T]) Push(value *T) {
	n := &node[T]{value, s.top}
	s.top = n
	s.length++
}
