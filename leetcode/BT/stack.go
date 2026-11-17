
type Stack struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() bool {

	top := s.items[len(s.items)-1]
	s.items = s.items[:len(items)-1]
	return top, true

}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return false
	}

	return true
}
