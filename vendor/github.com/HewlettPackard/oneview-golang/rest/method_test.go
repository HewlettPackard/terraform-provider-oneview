package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethod(t *testing.T) {
	var m Method
	m = GET
	assert.Equal(t, "GET", m.String(), "GET should be string")
	assert.Equal(t, "POST", POST.String(), "POST should be string")
	assert.Equal(t, "PUT", PUT.String(), "PUT should be string")
	assert.Equal(t, "DELETE", DELETE.String(), "DELETE should be string")
	assert.Equal(t, "PATCH", PATCH.String(), "PATCH should be string")
}
