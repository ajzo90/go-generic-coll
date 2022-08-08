package collection

type void struct {
}

type Set[K comparable] map[K]void

func (s Set[K]) Elements() []K {
	return MapKeys(s)
}

func (s Set[K]) Len() int {
	return len(s)
}

func (s Set[K]) Set(k K) {
	s[k] = void{}
}

func (s Set[K]) Has(k K) bool {
	_, has := s[k]
	return has
}

func (s Set[K]) Unset(k K) {
	delete(s, k)
}
