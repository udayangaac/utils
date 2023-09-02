// Package operator provides common operators for various types.
package operator

// Or returns the first non-nil value between a and b.
// If a is non-nil, it is returned; otherwise, b is returned.
// It works with any type specified by the type parameter T.
func Or[T any](a, b *T) T {
	var t T

	if a != nil {
		return *a
	}

	if b != nil {
		return *b
	}
	return t
}
