package variables

import (
	"testflowkit/internal/step_definitions/core/stepbuilder"
)

type steps struct{}

func GetAllSteps() []stepbuilder.Step {
	st := steps{}
	return []stepbuilder.Step{
		st.storeJSONPathIntoVariable(),
		st.storeElementContentIntoVariable(),
		st.storeCustomVariable(),
	}
}
