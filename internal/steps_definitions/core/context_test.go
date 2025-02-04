package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldInstanciateCorrectlyNewFrontendContext(t *testing.T) {
	timeout := "15s"
	headlessMode := true
	slowMotion := 10 * time.Millisecond

	frontendCtx := NewFrontendContext(timeout, headlessMode, slowMotion)

	assert.Equal(t, timeout, frontendCtx.timeout.String())
	assert.True(t, frontendCtx.headlessMode)
	assert.Equal(t, slowMotion, frontendCtx.slowMotion)
	assert.Nil(t, frontendCtx.browser)
	assert.Nil(t, frontendCtx.page)
}
