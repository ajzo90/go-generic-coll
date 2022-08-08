package collection

type Stack[V any] struct {
	v []V
}

func (s *Stack[V]) Push(v V) {
	s.v = append(s.v, v)
}

func (s *Stack[V]) Elements() []V {
	return s.v
}

func (s *Stack[V]) Len() int {
	return len(s.v)
}

func (s *Stack[V]) Peek() (V, bool) {
	if s.Len() == 0 {
		return *new(V), false
	}
	return s.v[len(s.v)-1], true
}

func (s *Stack[V]) Clear() {
	s.v = s.v[:0]
}

func (s *Stack[V]) Pop() (V, bool) {
	v, ok := s.Peek()
	if !ok {
		return v, ok
	}
	s.v = s.v[:len(s.v)-1]
	return v, ok
}
