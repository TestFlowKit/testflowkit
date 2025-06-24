package scenario

import (
	"testflowkit/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldInstanciateCorrectlyNewFrontendContext(t *testing.T) {
	cfg := config.Config{
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
	}
	ctx := NewContext(&cfg)

	assert.InDelta(t, float64(15), ctx.timeout.Seconds(), 0.0001)
	assert.False(t, ctx.headlessMode)
	assert.InDelta(t, float64(10), ctx.slowMotion.Seconds(), 0.0001)
	assert.Nil(t, ctx.browser)
	assert.Nil(t, ctx.page)
}
