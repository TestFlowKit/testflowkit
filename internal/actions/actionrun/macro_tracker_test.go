package actionrun

import (
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMacroTracker_CollapsesMacroGroups(t *testing.T) {
	mt := newMacroTracker([]macroGroup{{callText: "When a macro step", startIdx: 0, stepCount: 2}})

	out1 := mt.processStep("When a macro step 1", godog.StepPassed, time.Millisecond, nil)
	assert.False(t, out1.emit)

	out2 := mt.processStep("And a macro step 2", godog.StepPassed, time.Millisecond, nil)
	require.True(t, out2.emit)
	assert.Equal(t, "When a macro step", out2.title)
	assert.Equal(t, godog.StepPassed, out2.status)
	assert.NotZero(t, out2.dur)
}

func TestMacroTracker_RespectsAbsoluteStepOffsets(t *testing.T) {
	mt := newMacroTracker([]macroGroup{{callText: "When a macro step", startIdx: 1, stepCount: 2}})

	first := mt.processStep("Given a step", godog.StepPassed, 2*time.Millisecond, nil)
	assert.True(t, first.emit)
	assert.Equal(t, "Given a step", first.title)

	second := mt.processStep("When a macro step 1", godog.StepPassed, 3*time.Millisecond, nil)
	assert.False(t, second.emit)

	third := mt.processStep("And a macro step 2", godog.StepFailed, 4*time.Millisecond, assert.AnError)
	require.True(t, third.emit)
	assert.Equal(t, "When a macro step", third.title)
	assert.Equal(t, godog.StepFailed, third.status)
	assert.Equal(t, assert.AnError, third.err)
}
