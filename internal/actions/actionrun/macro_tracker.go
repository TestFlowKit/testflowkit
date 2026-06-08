package actionrun

import (
	"time"

	"testflowkit/pkg/gherkinparser"

	"github.com/cucumber/godog"
)

// macroGroup describes where a macro call was expanded within the combined
// (background + scenario) step sequence for a single scenario execution.
// startIdx is the absolute 0-based index of the first expanded step across
// the background + scenario step sequence.
type macroGroup struct {
	callText  string
	startIdx  int
	stepCount int
}

// macroTracker collapses expanded macro steps back into a single report step
// when the runner is in implicit report mode.
//
// Godog fires one step-hook call per expanded step. The tracker counts calls,
// detects when a group starts (startIdx match), accumulates status/duration/error
// across the group, and emits one synthetic step when the last call of the group
// arrives. Non-group steps are passed through unchanged.
type macroTracker struct {
	groups   []macroGroup
	stepIdx  int
	groupPtr int

	// in-group accumulation
	inGroup     bool
	stepsLeft   int
	accumDur    time.Duration
	accumStatus godog.StepResultStatus
	accumErr    error
	accumCall   string
}

// stepOutput is the decision record returned by processStep.
// When emit is false the step was absorbed into a macro group; callers must not
// add it to the report.
type stepOutput struct {
	emit   bool
	title  string
	status godog.StepResultStatus
	dur    time.Duration
	err    error
}

func newMacroTracker(groups []macroGroup) *macroTracker {
	return &macroTracker{
		groups:      groups,
		accumStatus: godog.StepPassed,
	}
}

// processStep processes one incoming step hook call. It returns a stepOutput
// that either passes the step through, absorbs it into the current macro group,
// or flushes the completed macro group as a single synthetic step.
func (mt *macroTracker) processStep(
	title string, status godog.StepResultStatus, dur time.Duration, err error,
) stepOutput {
	// Check whether this step starts a new macro group.
	if !mt.inGroup && mt.groupPtr < len(mt.groups) && mt.stepIdx == mt.groups[mt.groupPtr].startIdx {
		g := mt.groups[mt.groupPtr]
		mt.inGroup = true
		mt.stepsLeft = g.stepCount
		mt.accumDur = 0
		mt.accumStatus = godog.StepPassed
		mt.accumErr = nil
		mt.accumCall = g.callText
	}

	mt.stepIdx++

	if mt.inGroup {
		mt.accumDur += dur
		mt.accumStatus = worseStatus(mt.accumStatus, status)
		if err != nil && mt.accumErr == nil {
			mt.accumErr = err
		}
		mt.stepsLeft--

		if mt.stepsLeft == 0 {
			// Group complete — flush as one synthetic step.
			callText := mt.accumCall
			flushedStatus := mt.accumStatus
			flushedDur := mt.accumDur
			flushedErr := mt.accumErr
			mt.inGroup = false
			mt.groupPtr++
			return stepOutput{
				emit:   true,
				title:  callText,
				status: flushedStatus,
				dur:    flushedDur,
				err:    flushedErr,
			}
		}
		// Still inside the group — absorb this step.
		return stepOutput{emit: false}
	}

	// Regular step — pass through.
	return stepOutput{
		emit:   true,
		title:  title,
		status: status,
		dur:    dur,
		err:    err,
	}
}

// stepStatusPriority returns the relative severity of a step result status.
// Priority (ascending): Passed < Pending < Undefined < Skipped < Ambiguous < Failed.
func stepStatusPriority(status godog.StepResultStatus) int {
	switch status {
	case godog.StepPassed:
		return 0
	case godog.StepPending:
		return 1
	case godog.StepUndefined:
		return 2
	case godog.StepSkipped:
		return 3
	case godog.StepAmbiguous:
		return 4
	case godog.StepFailed:
		return 5
	default:
		return 0
	}
}

// worseStatus returns whichever of a or b represents a worse outcome.
// Priority (ascending): Passed < Pending < Undefined < Skipped < Ambiguous < Failed.
func worseStatus(a, b godog.StepResultStatus) godog.StepResultStatus {
	if stepStatusPriority(b) > stepStatusPriority(a) {
		return b
	}
	return a
}

// buildMacroGroups converts the macro expansion metadata from a Feature into a
// sorted slice of macroGroup values with absolute step indices (background steps
// come first, then scenario steps offset by BackgroundStepCount).
func buildMacroGroups(f *gherkinparser.Feature, scenarioName string) []macroGroup {
	var groups []macroGroup

	for _, entry := range f.BackgroundMacros {
		groups = append(groups, macroGroup{
			callText:  entry.CallText,
			startIdx:  entry.StartIdx,
			stepCount: entry.StepCount,
		})
	}

	for _, entry := range f.ScenarioMacros[scenarioName] {
		groups = append(groups, macroGroup{
			callText:  entry.CallText,
			startIdx:  f.BackgroundStepCount + entry.StartIdx,
			stepCount: entry.StepCount,
		})
	}

	// Background entries are already ordered first; scenario entries are appended
	// after. A sort is not strictly necessary given parse order, but guards against
	// any future reordering.
	for i := 1; i < len(groups); i++ {
		for j := i; j > 0 && groups[j].startIdx < groups[j-1].startIdx; j-- {
			groups[j], groups[j-1] = groups[j-1], groups[j]
		}
	}

	return groups
}
