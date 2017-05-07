package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainMySum(t *testing.T) {
	assert.Equal(t, 3, mySum(1, 2))
	assert.Equal(t, 0, mySum(1, -1))
}

func TestMainMyPosSub(t *testing.T) {
	z, err := myPosSub(2, 1)
	assert.NoError(t, err)
	assert.Equal(t, 1, z)
	z, err = myPosSub(1, 1)
	assert.Error(t, err)
}

func TestMainMyMultiply(t *testing.T) {
	assert.Equal(t, 2, myMultiply(1, 2))
	assert.Equal(t, 1, myMultiply(1, 1))
	assert.Equal(t, 1, myMultiply(-1, -1))
	assert.Equal(t, -1, myMultiply(-1, 1))
}
