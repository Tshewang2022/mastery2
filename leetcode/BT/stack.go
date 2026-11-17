package main

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, bool) {

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true

}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.items) == 0 {
		return false
	}

	return true
}
