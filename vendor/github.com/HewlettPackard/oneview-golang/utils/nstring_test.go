package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNstring(t *testing.T) {
	assert.Equal(t, Nstring("foo"), NewNstring("foo"), "create Nstring should be equal")
	var n Nstring
	n.Nil()
	assert.Equal(t, "null", n.String(), "should create an empty null string")
	n = NewNstring("null")
	assert.Equal(t, "null", n.String(), "should create an empty null string")
	n = "foobar"
	assert.Equal(t, "foobar", n.String(), "should container foobar")
}

func TestIsNil(t *testing.T) {
	var n Nstring
	assert.True(t, n.IsNil(), "should be a nil string")
}
