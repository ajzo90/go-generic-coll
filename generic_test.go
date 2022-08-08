package collection

import (
	"fmt"
	"log"
	"testing"
)

func Test_Stuff(t *testing.T) {
	type Record struct {
		Id   int    `json:"id"`
		Type string `json:"type"`
	}
	var sl = []Record{{Id: 1, Type: "A"}, {Id: 2, Type: "A"}}

	var res = SliceFilter(sl, []Record{}, func(v Record, idx int) bool {
		return v.Id == 1
	})
	log.Println(res)

	var a = SliceMap(sl, []Record{}, func(from Record, idx int) Record {
		return from
	})
	log.Println(a)

	var acc = SliceReduce(sl, Set[int]{}, func(acc Set[int], v Record, idx int) Set[int] {
		acc.Set(v.Id)
		return acc
	})
	log.Println(acc)

	var m = map[string]Record{"xasd": {Type: "A"}, "xasda": {Type: "X"}}
	log.Println(MapKeys(m))
	log.Println(MapValues(m))

	k := SliceMap(MapKeys(m), []string{}, AppendV("xxx"))
	log.Println(k)
	log.Println(typeOf(Record{}))

	log.Println(MapMap(m, []string{}, func(V Record, k string) string {
		return k + V.Type
	}))

	log.Println(MapDistinct(m, map[string]Record{}, func(v Record, k string) string {
		return ""
	}))

	log.Println(MapEntries(m))

	log.Println(Map(Slice(sl).It(), []int{}, func(r Record, k int) int {
		return k
	}))

	log.Println(Map(Slice(sl).PtrIt(), []string{}, func(r *Record, k int) string {
		return r.Type
	}))

	Slice(sl).ToStrings(func(v Record, idx int) string {
		return v.Type
	})

}

type Addable interface {
	Ordered
}

func AppendV[T Addable](s T) func(v T, i int) T {
	return func(v T, i int) T {
		return v + s
	}
}

func typeOf[T any](v T) string {
	return fmt.Sprintf("%T", *new(T))
}
