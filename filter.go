package clise

// FilterFunc returns true if an element is to be kept in a slice.
// This function must not have side effects on the target slice.
type FilterFunc func(index int) bool
