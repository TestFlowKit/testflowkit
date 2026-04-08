package commonbackendsteps

import (
	"errors"
	"strings"

	"testflowkit/internal/step_definitions/core/scenario"
	"testflowkit/pkg/queryable"
)

func newResponseEngine(backend *scenario.BackendContext) (queryable.Queryable, error) {
	if backend == nil || !backend.HasResponse() {
		return nil, errors.New("no response available")
	}

	responseBody := backend.GetResponseBody()
	if responseBody == nil {
		return nil, errors.New("response body is empty")
	}

	return queryable.NewEngine(responseBody, detectResponseFormat(backend.GetResponseHeaders()))
}

func detectResponseFormat(headers map[string]string) queryable.Format {
	for key, value := range headers {
		if strings.EqualFold(key, "Content-Type") {
			return queryable.DetectFormatFromContentType(value)
		}
	}
	return queryable.FormatAuto
}
