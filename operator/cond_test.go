// Package operator_test contains test cases for the operator package.
package operator_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udayangaac/utils/operator"
)

// TestOr tests the Or function from the operator package.
func TestOr(t *testing.T) {
	// Create an expected error.
	var expectedErr error = errors.New("sample error")

	// Test Or with a nil error and an expected error.
	actualErr := operator.Or(nil, &expectedErr)
	assert.Equal(t, expectedErr, actualErr)

	// Test Or with an expected error and a nil error.
	actualErr = operator.Or(&expectedErr, nil)
	assert.Equal(t, expectedErr, actualErr)
}
