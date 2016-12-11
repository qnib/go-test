package main

import (
	"testing"

    "github.com/stretchr/testify/assert"
)

func TestStupidMin(t *testing.T) {
    assert.Equal(t, 2, stupidMin(2, 3))
    assert.Equal(t, 2, stupidMin(3, 2))
}
