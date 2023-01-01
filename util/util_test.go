//go:build unittest
// +build unittest

package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_RandomInt(t *testing.T) {
	randomNumber := RandomInt(10, 100)

	require.NotEmpty(t, randomNumber)

	checker := false
	if (randomNumber >= 10) && (randomNumber <= 100) {
		checker = true
	}

	require.Equal(t, true, checker)
}

func Test_RandomString(t *testing.T) {
	randomString := RandomString(100)

	require.NotEmpty(t, randomString)
	require.Equal(t, 100, len(randomString))
}
