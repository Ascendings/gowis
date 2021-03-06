package base

// Append ...
func Append(slice []string, items ...string) []string {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

// Extend ...
func Extend(slice []string, element string) []string {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]string, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
