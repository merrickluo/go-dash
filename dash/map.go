package dash

// Keys returns a slice contains all keys in map m
func Keys[K comparable, V any](m map[K]V) []K {
	ret := make([]K, 0, len(m))

	for k := range m {
		ret = append(ret, k)
	}

	return ret
}

// Keys returns a slice contains all values in map m
func Values[K comparable, V any](m map[K]V) []V {
	ret := make([]V, 0, len(m))

	for _, v := range m {
		ret = append(ret, v)
	}

	return ret
}

// Merge maps into one map, uses the value in the latter map.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	return MergeWith(func(v1 V, v2 V) V {
		return v2
	}, maps...)
}

// MergeWith returns a new map contains all keys in maps
// and values by applying function f to all values with the same key
func MergeWith[K comparable, V any](f func(v1 V, v2 V) V, maps ...map[K]V) map[K]V {
	ret := make(map[K]V)
	for _, m := range maps {
		for k, v := range m {
			if rv, found := ret[k]; found {
				ret[k] = f(rv, v)
			} else {
				ret[k] = v
			}
		}
	}

	return ret
}
