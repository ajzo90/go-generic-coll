package collection

func firstOrNew[T any, Coll ~[]T](prealloc ...Coll) Coll {
	if len(prealloc) == 0 {
		return *new(Coll)
	} else {
		return prealloc[0]
	}
}
