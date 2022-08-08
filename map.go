package collection

func MapReduce[K comparable, V any, M ~map[K]V, Acc any](m M, acc Acc, f func(Acc, V, K) Acc) Acc {
	for k, v := range m {
		acc = f(acc, v, k)
	}
	return acc
}

func MapMap[K comparable, V any, M ~map[K]V, T any, Coll ~[]T](m M, dst Coll, f func(v V, k K) T) Coll {
	return MapReduce(m, Grow(dst, len(m)), func(coll Coll, v V, k K) Coll {
		return append(coll, f(v, k))
	})
}

func MapForeach[K comparable, V any, M ~map[K]V](m M, f func(v V, k K)) {
	for k, v := range m {
		f(v, k)
	}
}

func MapKeys[K comparable, V any, M ~map[K]V](m M, prealloc ...[]K) []K {
	return MapMap(m, firstOrNew(prealloc...), func(v V, k K) K {
		return k
	})
}

func MapValues[K comparable, V any, M ~map[K]V](m M, prealloc ...[]V) []V {
	return MapMap(m, firstOrNew(prealloc...), func(v V, k K) V {
		return v
	})
}

func MapFilter[K comparable, V any, M ~map[K]V](src, dest M, pred func(V, K) bool) M {
	return MapReduce(src, dest, func(dst M, v V, k K) M {
		if pred(v, k) {
			dst[k] = v
		} else {
			delete(dst, k)
		}
		return dst
	})
}

func MapFind[K comparable, V any, M ~map[K]V](src M, pred func(V, K) bool) (K, V, bool) {
	for k, v := range src {
		if pred(v, k) {
			return k, v, true
		}
	}
	return *new(K), *new(V), false
}

func MapExist[K comparable, V any, M ~map[K]V](src M, pred func(V, K) bool) bool {
	_, _, ok := MapFind(src, pred)
	return ok
}

func MapDistinct[K comparable, V any, M ~map[K]V, T comparable](src, dest M, f func(V, K) T) M {
	var seen = Set[T]{}
	return MapFilter(src, dest, func(v V, k K) bool {
		key := f(v, k)
		if _, exist := seen[key]; !exist {
			seen.Set(key)
			return true
		}
		return false
	})
}

type MapEntry[K comparable, V any] struct {
	k K
	v V
}

func MapEntries[K comparable, V any, M ~map[K]V, E MapEntry[K, V]](m M) []E {
	return MapMap(m, []E{}, func(v V, k K) E {
		return E{k: k, v: v}
	})
}
