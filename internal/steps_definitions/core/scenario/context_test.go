package scenario

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldInstanciateCorrectlyNewFrontendContext(t *testing.T) {
	timeout := "15s"
	headlessMode := true
	slowMotion := 10 * time.Millisecond

	ctx := NewContext(timeout, headlessMode, slowMotion)

	assert.Equal(t, timeout, ctx.timeout.String())
	assert.True(t, ctx.headlessMode)
	assert.Equal(t, slowMotion, ctx.slowMotion)
	assert.Nil(t, ctx.browser)
	assert.Nil(t, ctx.page)
}
