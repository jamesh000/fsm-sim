package stack

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return top
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
