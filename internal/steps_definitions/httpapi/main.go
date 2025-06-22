package httpapi

import (
	"testflowkit/internal/config"
	"testflowkit/internal/steps_definitions/core/stepbuilder"
	"testflowkit/pkg/logger"
)

type steps struct {
}

// GetAllSteps returns all available HTTP API step definitions.
// This function automatically detects whether enhanced configuration is available
// and returns the appropriate step definitions accordingly.
//
// The function provides:
// - Legacy steps for backward compatibility
// - Enhanced steps when enhanced configuration is available
// - Comprehensive API testing capabilities
func GetAllSteps() []stepbuilder.Step {
	st := steps{}
	
	// Base legacy steps for backward compatibility
	allSteps := []stepbuilder.Step{
		st.prepareRequest(),
		st.setHeaders(),
		st.setQueryParams(),
		st.setRequestBody(),
		st.setPathParams(),
		st.sendRequest(),
		st.checkResponseStatusCode(),
		st.responseBodyShouldContain(),
		st.responseBodyPathShouldExist(),
		st.storeResponseBodyPath(),
		st.storeResponseHeader(),
	}

	// Add enhanced steps if enhanced configuration is available
	if isEnhancedConfigurationAvailable() {
		logger.Info("Enhanced configuration detected - registering enhanced API steps")
		enhancedSteps := []stepbuilder.Step{
			st.prepareRequestEnhanced(),
			st.setRequestHeaderEnhanced(),
			st.setPathParameterEnhanced(),
			st.setQueryParameterEnhanced(),
			st.setRequestBodyEnhanced(),
			st.sendRequestEnhanced(),
			st.verifyResponseStatusEnhanced(),
			st.verifyResponseContainsEnhanced(),
			st.verifyResponseHeaderEnhanced(),
			st.storeResponseValueEnhanced(),
		}
		allSteps = append(allSteps, enhancedSteps...)
	} else {
		logger.Info("Legacy configuration detected - using legacy API steps only")
	}

	return allSteps
}

// isEnhancedConfigurationAvailable checks if enhanced configuration is loaded and available.
// This allows for graceful fallback to legacy steps when enhanced configuration is not used.
func isEnhancedConfigurationAvailable() bool {
	_, err := config.GetEnhancedConfig()
	return err == nil
}
