package collection

func Grow[S ~[]E, E any](s S, n int) S {
	return append(s, make(S, n)...)[:len(s)]
}

func SliceMap[Src any, Dst any, SrcColl ~[]Src, DstColl ~[]Dst](s SrcColl, res DstColl, fn func(Src, int) Dst) DstColl {
	res = Grow(res, len(s))
	for i, v := range s {
		res = append(res, fn(v, i))
	}
	return res
}

func SliceToStrings[Src any, SrcColl ~[]Src](s SrcColl, fn func(Src, int) string) []string {
	return SliceMap(s, []string{}, fn)
}

func SliceToInts[Src any, SrcColl ~[]Src](s SrcColl, fn func(Src, int) int) []int {
	return SliceMap(s, []int{}, fn)
}

func SliceToSlice[Src any, SrcColl ~[]Src](s SrcColl, fn func(Src, int) Src) SrcColl {
	return SliceMap(s, SrcColl{}, fn)
}

func SliceReduce[T any, Coll ~[]T, Acc any](coll Coll, acc Acc, f func(Acc, T, int) Acc) Acc {
	for i, v := range coll {
		acc = f(acc, v, i)
	}
	return acc
}

func SliceFilter[T any, Coll ~[]T](src, dest Coll, fn func(T, int) bool) Coll {
	for i, v := range src {
		if fn(v, i) {
			dest = append(dest, v)
		}
	}
	return dest
}

func SliceUniq[T comparable, Coll ~[]T](src, dest Coll) Coll {
Loop:
	for _, v := range src {
		for _, v2 := range dest {
			if v == v2 {
				continue Loop
			}
		}
		dest = append(dest, v)
	}
	return dest
}

func SliceCount[T any, Coll ~[]T](coll Coll, fn func(T, int) bool) int {
	var count int
	for i, v := range coll {
		if fn(v, i) {
			count++
		}
	}
	return count
}
