package collection

import "sort"

func Slice[V any, Coll ~[]V](items Coll) SliceT[V, Coll] {
	return SliceT[V, Coll](items)
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

type SliceT[V any, Coll ~[]V] []V

func (s SliceT[V, Coll]) Iterate(f func(V, int)) {
	for i, v := range s {
		f(v, i)
	}
}

func (s SliceT[V, Coll]) It() Iterator[V, int] {
	return s
}

func (s SliceT[V, Coll]) PtrIt() Iterator[*V, int] {
	return SlicePtrT[V, Coll](s)
}

type SlicePtrT[V any, Coll ~[]V] []V

func (s SlicePtrT[V, Coll]) Len() int {
	return len(s)
}

func (s SlicePtrT[V, Coll]) Iterate(f func(*V, int)) {
	for i := range s {
		f(&s[i], i)
	}
}

func (s SliceT[V, Coll]) Filter(f func(V, int) bool) SliceT[V, Coll] {
	return SliceFilter(s, SliceT[V, Coll]{}, f)
}

func (s SliceT[V, Coll]) Find(f func(V, int) bool) (V, int) {
	for i, v := range s {
		if f(v, i) {
			return v, i
		}
	}
	return *new(V), -1
}

func (s SliceT[V, Coll]) Count(f func(V, int) bool) int {
	var count int
	for i, v := range s {
		if f(v, i) {
			count++
		}
	}
	return count
}

func (s SliceT[V, Coll]) Sorted(f func(a, b V) bool) SliceT[V, Coll] {
	dest := append(SliceT[V, Coll]{}, s...)
	sort.Slice(dest, func(i, j int) bool {
		return f(dest[i], dest[j])
	})
	return dest
}

func (s SliceT[V, Coll]) Len() int {
	return len(s)
}

func (s SliceT[V, Coll]) ToInts(f func(V, int) int) []int {
	return SliceToInts(s, f)
}

func (s SliceT[V, Coll]) ToStrings(f func(V, int) string) SliceT[string, []string] {
	return SliceToStrings(s, f)
}

func (s SliceT[V, Coll]) Map(f func(V, int) V) SliceT[V, Coll] {
	return SliceToSlice(s, f)
}
