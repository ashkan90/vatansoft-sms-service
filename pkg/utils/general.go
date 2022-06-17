package utils

import "sort"

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		string
}

func InArray[E comparable](s []E, v E) bool {
	return Index(s, v) != -1
}

// Index returns the index of the first occurrence of v in s, or -1 if not present.
func Index[E comparable](s []E, v E) int {
	for i, vs := range s {
		if v == vs {
			return i
		}
	}
	return -1
}

func SortedKeys[K Ordered, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}

func OrderedComparison[K Ordered, V any](m map[K]V, com K) any {
	var (
		keys    = SortedKeys(m)
		ln      = len(keys)
		next    = 1
		nextVal = K(0)
		currVal = K(0)
	)

	for i := range keys {
		if next < ln {
			nextVal = keys[next]
			currVal = keys[i]

			if com > currVal && com <= nextVal {
				return m[currVal]
			}
		}

		next += 1
	}

	return -1
}
