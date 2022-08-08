package collection

type MapIter[V any, K comparable] map[K]V

func (m MapIter[V, K]) Iterate(f func(V, K)) {
	for k, v := range m {
		f(v, k)
	}
}

func (m MapIter[V, K]) Len() int {
	return len(m)
}

type Iterator[V any, K comparable] interface {
	Iterate(func(V, K))
	Len() int
}

func Map[V any, K comparable, R any, DstColl ~[]R](it Iterator[V, K], dst DstColl, f func(V, K) R) []R {
	dst = Grow(dst, it.Len())
	it.Iterate(func(v V, k K) {
		dst = append(dst, f(v, k))
	})
	return dst
}

func ToSet[V comparable, K comparable](it Iterator[V, K]) Set[V] {
	var s = Set[V]{}
	it.Iterate(func(v V, k K) {
		s.Set(v)
	})
	return s
}
