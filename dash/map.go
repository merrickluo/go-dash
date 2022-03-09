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

// MergeWith returns a new map contains all keys in maps
// and values by applying function f to all values with the same key
func MergeWith[K comparable, V any, NV any](f func(...V) NV, maps ...map[K]V) map[K]NV {
	keys := Uniq(FlatMap(maps, Keys[K, V]))
	ret := make(map[K]NV, len(keys))

	for _, k := range keys {
		vs := make([]V, 0, len(keys))
		for _, m := range maps {
			if v, found := m[k]; found {
				vs = append(vs, v)
			}
		}

		ret[k] = f(vs...)
	}

	return ret
}
