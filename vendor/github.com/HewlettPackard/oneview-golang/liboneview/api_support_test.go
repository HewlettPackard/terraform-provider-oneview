package liboneview

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNew
func TestNew(t *testing.T) {
	var asc APISupport
	asc = asc.New(asc.Get("foo"))
	assert.True(t, (asc == C_NONE))
	assert.True(t, (asc.New(asc.Integer()) == C_NONE))
}

// TestHasCheck
func TestHasCheck(t *testing.T) {
	var asc APISupport
	assert.False(t, asc.HasCheck("foo"))
	assert.True(t, asc.HasCheck("profile_templates.go"))
}

// TestGet
func TestGet(t *testing.T) {
	var asc APISupport
	assert.Equal(t, int(C_NONE), asc.Get("foo"))
	assert.Equal(t, int(C_PROFILE_TEMPLATES), asc.Get("profile_templates.go"))
}

// TestCheck
func TestIsSupported(t *testing.T) {
	var currentversion Version
	var asc APISupport
	currentversion = API_VER1 // force current version to ver 1 for testing

	asc = asc.New(asc.Get("foo"))
	assert.True(t, (asc == C_NONE))                 // is same as C_NONE
	assert.True(t, asc.IsSupported(currentversion)) // should be true for any version

	asc = asc.New(asc.Get("profile_templates.go"))
	assert.True(t, (asc == C_PROFILE_TEMPLATES))     // is same as C_PROFILE_TEMPLATES
	assert.False(t, asc.IsSupported(currentversion)) // should be false for ver1
	assert.True(t, asc.IsSupported(API_VER2))        // should be true for ver2
	assert.True(t, asc.IsSupported(API_VER_UNKNOWN)) // should still be supported just unkown

	// simulate client versions
	var clientversion Version
	clientversion = clientversion.CalculateVersion(120, 108)
	asc = asc.NewByName("profile_templates.go")
	assert.False(t, asc.IsSupported(clientversion)) // should not be supported
}

// TestNewByName - get a new APISupport object by name
func TestNewByName(t *testing.T) {
	var asc APISupport
	asc = asc.NewByName("foo")
	assert.True(t, (asc == C_NONE))
	assert.False(t, (asc == C_PROFILE_TEMPLATES))

	asc = asc.NewByName("profile_templates.go")
	assert.False(t, (asc == C_NONE))
	assert.True(t, (asc == C_PROFILE_TEMPLATES))
}
