package utils

func CompareTwoArraysByIntKey[T1, T2 any](a []T1, b []T2, getKey1 func(T1) uint, getKey2 func(T2) uint) bool {
	if len(a) != len(b) {
		return false
	}

	// Count frequency of IDs in slice A
	freq := make(map[uint]uint)
	for _, item := range a {
		key := getKey1(item)
		freq[key]++
	}

	// Compare with slice B
	for _, item := range b {
		key := getKey2(item)
		if freq[key] == 0 {
			return false // ID not found or count mismatched
		}
		freq[key]--
	}

	return true
}
