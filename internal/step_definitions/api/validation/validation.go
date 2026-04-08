package validation

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/antchfx/xmlquery"

	"testflowkit/internal/step_definitions/api/jsonhelpers"
	"testflowkit/pkg/apperrors"
	"testflowkit/pkg/logger"
	"testflowkit/pkg/queryable"
)

const NoPathValueFound = "failed to get value at path '%s': %w"

// ValidatePathValue validates that a response path contains or does not contain
// the expected value, regardless of whether the response is JSON or XML.
func ValidatePathValue(engine queryable.Queryable, path, expectedValue string, shouldEqual bool) error {
	if engine == nil {
		return apperrors.ErrNoResponseAvailable
	}

	result, err := engine.Get(path)
	if err != nil {
		return fmt.Errorf(NoPathValueFound, path, err)
	}
	if !result.Exists {
		return fmt.Errorf("path '%s' does not exist in response", path)
	}

	if shouldEqual && result.Raw != expectedValue {
		return fmt.Errorf("expected value '%s' at path '%s', but got '%s'", expectedValue, path, result.Raw)
	}

	if !shouldEqual && result.Raw == expectedValue {
		return fmt.Errorf("field '%s' has forbidden value '%s'", path, result.Raw)
	}

	if shouldEqual {
		logger.InfoFf("response validation passed: %s = %s", path, expectedValue)
	} else {
		logger.InfoFf("response validation passed: %s != %s", path, expectedValue)
	}
	return nil
}

// ValidatePathExists validates that a response path exists or does not exist.
func ValidatePathExists(engine queryable.Queryable, path string, shouldExist bool) error {
	if engine == nil {
		return apperrors.ErrNoResponseAvailable
	}

	exists, err := engine.Exists(path)
	if err != nil {
		return err
	}

	if shouldExist && !exists {
		return fmt.Errorf("path '%s' does not exist in response", path)
	}

	if !shouldExist && exists {
		return fmt.Errorf("path '%s' exists in response but should not", path)
	}

	if shouldExist {
		logger.InfoFf("response validation passed: path '%s' exists", path)
	} else {
		logger.InfoFf("response validation passed: path '%s' does not exist", path)
	}
	return nil
}

// ValidatePathContains validates substring presence on a single response field value.
func ValidatePathContains(engine queryable.Queryable, path, expectedSubstring string, shouldContain bool) error {
	if engine == nil {
		return apperrors.ErrNoResponseAvailable
	}

	result, err := engine.Get(path)
	if err != nil {
		return fmt.Errorf(NoPathValueFound, path, err)
	}
	if !result.Exists {
		return fmt.Errorf("path '%s' does not exist in response", path)
	}

	contains := strings.Contains(result.Raw, expectedSubstring)
	if shouldContain && !contains {
		return fmt.Errorf("field '%s' value '%s' does not contain expected text '%s'", path, result.Raw, expectedSubstring)
	}

	if !shouldContain && contains {
		return fmt.Errorf("field '%s' value '%s' should not contain text '%s'", path, result.Raw, expectedSubstring)
	}

	return nil
}

// ValidatePathPattern validates regex matching on a single response field value.
func ValidatePathPattern(engine queryable.Queryable, path, pattern string, shouldMatch bool) error {
	if engine == nil {
		return apperrors.ErrNoResponseAvailable
	}

	result, err := engine.Get(path)
	if err != nil {
		return fmt.Errorf(NoPathValueFound, path, err)
	}
	if !result.Exists {
		return fmt.Errorf("path '%s' does not exist in response", path)
	}

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("invalid regex pattern '%s': %w", pattern, err)
	}

	matches := regex.MatchString(result.Raw)
	if shouldMatch && !matches {
		return fmt.Errorf("field '%s' value '%s' does not match pattern '%s'", path, result.Raw, pattern)
	}

	if !shouldMatch && matches {
		return fmt.Errorf("field '%s' value '%s' should not match pattern '%s'", path, result.Raw, pattern)
	}

	return nil
}

// ValidatePathType validates the normalized type of the data found at a path.
func ValidatePathType(engine queryable.Queryable, path, expectedType string, shouldMatchType bool) error {
	if engine == nil {
		return apperrors.ErrNoResponseAvailable
	}

	actualType, err := resolvePathType(engine, path)
	if err != nil {
		return err
	}

	normalizedExpectedType := queryable.NormalizeType(expectedType)
	if shouldMatchType && actualType != normalizedExpectedType {
		return fmt.Errorf("field '%s' has type '%s', expected '%s'", path, actualType, normalizedExpectedType)
	}

	if !shouldMatchType && actualType == normalizedExpectedType {
		return fmt.Errorf("field '%s' has forbidden type '%s'", path, actualType)
	}

	return nil
}

// ValidatePathCount validates the number of elements returned for a path.
func ValidatePathCount(engine queryable.Queryable, path string, expectedCount int) error {
	if engine == nil {
		return apperrors.ErrNoResponseAvailable
	}

	results, err := engine.GetAll(path)
	if err != nil {
		return fmt.Errorf("failed to get values at path '%s': %w", path, err)
	}

	actualCount := len(results)
	if actualCount != expectedCount {
		return fmt.Errorf("expected %d elements in '%s', but got %d", expectedCount, path, actualCount)
	}

	logger.InfoFf("response validation passed: path '%s' contains %d element(s)", path, expectedCount)
	return nil
}

func resolvePathType(engine queryable.Queryable, path string) (string, error) {
	results, err := engine.GetAll(path)
	if err != nil {
		return "", fmt.Errorf(NoPathValueFound, path, err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf(NoPathValueFound, path, err)
	}
	if len(results) > 1 {
		return string(queryable.TypeArray), nil
	}
	if results[0].Kind != "" {
		return queryable.NormalizeType(results[0].Kind), nil
	}

	return queryable.GetValueType(results[0].Value), nil
}

// ValidateJSONPathValue validates that a JSON path contains the expected value.
func ValidateJSONPathValue(jsonBody []byte, jsonPath, expectedValue string) error {
	engine, err := queryable.NewEngine(jsonBody, queryable.FormatJSON)
	if err != nil {
		return err
	}
	return ValidatePathValue(engine, jsonPath, expectedValue, true)
}

// ValidateJSONPathExists validates that a JSON path exists in the response.
func ValidateJSONPathExists(jsonBody []byte, jsonPath string) error {
	engine, err := queryable.NewEngine(jsonBody, queryable.FormatJSON)
	if err != nil {
		return err
	}
	return ValidatePathExists(engine, jsonPath, true)
}

// ValidateBodyContains validates that the response body contains the expected substring.
func ValidateBodyContains(body []byte, expectedSubstring string) error {
	if body == nil {
		return apperrors.ErrNoResponseAvailable
	}

	bodyStr := string(body)
	if !strings.Contains(bodyStr, expectedSubstring) {
		return fmt.Errorf("response body does not contain expected substring '%s'", expectedSubstring)
	}

	logger.InfoFf("response validation passed: body contains '%s'", expectedSubstring)
	return nil
}

// ValidateBodyEquals validates that the response body matches the expected JSON or XML.
func ValidateBodyEquals(actual []byte, expected string) error {
	if actual == nil {
		return apperrors.ErrNoResponseAvailable
	}

	actualBody := bytes.TrimSpace(actual)
	expectedBody := bytes.TrimSpace([]byte(expected))
	if len(actualBody) == 0 {
		return errors.New("response body is empty")
	}
	if len(expectedBody) == 0 {
		return errors.New("expected body is empty")
	}

	actualFormat := detectBodyFormat(actualBody)
	expectedFormat := detectBodyFormat(expectedBody)
	if actualFormat != expectedFormat {
		return fmt.Errorf(
			"response body format mismatch: actual is %s but expected is %s",
			actualFormat,
			expectedFormat,
		)
	}

	switch actualFormat {
	case queryable.FormatXML:
		if err := compareXMLBodies(expectedBody, actualBody); err != nil {
			return fmt.Errorf("response XML validation failed: %w", err)
		}
		logger.InfoFf("response validation passed: XML matches expected content")
	case queryable.FormatJSON, queryable.FormatAuto:
		if err := jsonhelpers.CompareJSON(expectedBody, actualBody); err != nil {
			return fmt.Errorf("response JSON validation failed: %w", err)
		}
		logger.InfoFf("response validation passed: JSON matches expected content")
	}

	return nil
}

// ValidateJSONBodyEquals validates that the response body matches the expected JSON.
func ValidateJSONBodyEquals(actual []byte, expected string) error {
	return ValidateBodyEquals(actual, expected)
}

func detectBodyFormat(body []byte) queryable.Format {
	trimmedBody := bytes.TrimSpace(body)
	if len(trimmedBody) == 0 {
		return queryable.FormatAuto
	}
	if trimmedBody[0] == '<' {
		return queryable.FormatXML
	}
	return queryable.FormatJSON
}

func compareXMLBodies(expected, actual []byte) error {
	expectedDoc, err := xmlquery.Parse(bytes.NewReader(expected))
	if err != nil {
		return fmt.Errorf("failed to parse expected XML: %w", err)
	}

	actualDoc, err := xmlquery.Parse(bytes.NewReader(actual))
	if err != nil {
		return fmt.Errorf("failed to parse actual XML: %w", err)
	}

	expectedRoot := firstElementNode(expectedDoc)
	actualRoot := firstElementNode(actualDoc)
	if expectedRoot == nil || actualRoot == nil {
		return errors.New("xml document must contain a root element")
	}

	return compareXMLNodes(expectedRoot, actualRoot, "/"+expectedRoot.Data)
}

func firstElementNode(node *xmlquery.Node) *xmlquery.Node {
	if node == nil {
		return nil
	}
	if node.Type == xmlquery.ElementNode {
		return node
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if found := firstElementNode(child); found != nil {
			return found
		}
	}

	return nil
}

func compareXMLNodes(expected, actual *xmlquery.Node, path string) error {
	if err := validateXMLNodePair(expected, actual, path); err != nil {
		return err
	}

	if expected.Type == xmlquery.TextNode {
		return compareXMLText(path, expected.Data, actual.Data)
	}

	if expected.Type == xmlquery.ElementNode {
		if err := compareXMLElement(expected, actual, path); err != nil {
			return err
		}
	}

	return compareXMLChildren(expected, actual, path)
}

func validateXMLNodePair(expected, actual *xmlquery.Node, path string) error {
	if expected == nil || actual == nil {
		return fmt.Errorf("missing XML node at %s", path)
	}
	if expected.Type != actual.Type {
		return fmt.Errorf("node type mismatch at %s", path)
	}
	return nil
}

func compareXMLElement(expected, actual *xmlquery.Node, path string) error {
	if expected.Data != actual.Data {
		return fmt.Errorf("element mismatch at %s: expected <%s> but got <%s>", path, expected.Data, actual.Data)
	}
	return compareXMLAttributes(expected, actual, path)
}

func compareXMLChildren(expected, actual *xmlquery.Node, path string) error {
	expectedChildren := significantXMLChildren(expected)
	actualChildren := significantXMLChildren(actual)
	if len(expectedChildren) == 0 && len(actualChildren) == 0 {
		return compareXMLText(path, expected.InnerText(), actual.InnerText())
	}
	if len(expectedChildren) != len(actualChildren) {
		return fmt.Errorf(
			"child count mismatch at %s: expected %d but got %d",
			path,
			len(expectedChildren),
			len(actualChildren),
		)
	}

	for i := range expectedChildren {
		if err := compareXMLNodes(
			expectedChildren[i],
			actualChildren[i],
			xmlChildPath(path, expectedChildren[i]),
		); err != nil {
			return err
		}
	}

	return nil
}

func compareXMLText(path, expectedText, actualText string) error {
	normalizedExpected := strings.TrimSpace(expectedText)
	normalizedActual := strings.TrimSpace(actualText)
	if normalizedExpected != normalizedActual {
		return fmt.Errorf(
			"text mismatch at %s: expected %q but got %q",
			path,
			normalizedExpected,
			normalizedActual,
		)
	}
	return nil
}

func xmlChildPath(path string, node *xmlquery.Node) string {
	if node != nil && node.Type == xmlquery.ElementNode {
		return path + "/" + node.Data
	}
	return path
}

func compareXMLAttributes(expected, actual *xmlquery.Node, path string) error {
	if len(expected.Attr) != len(actual.Attr) {
		return fmt.Errorf(
			"attribute count mismatch at %s: expected %d but got %d",
			path,
			len(expected.Attr),
			len(actual.Attr),
		)
	}

	actualAttributes := make(map[string]string, len(actual.Attr))
	for _, attr := range actual.Attr {
		actualAttributes[attr.Name.Local] = attr.Value
	}

	for _, attr := range expected.Attr {
		actualValue, ok := actualAttributes[attr.Name.Local]
		if !ok {
			return fmt.Errorf("missing attribute %q at %s", attr.Name.Local, path)
		}
		if actualValue != attr.Value {
			return fmt.Errorf(
				"attribute %q mismatch at %s: expected %q but got %q",
				attr.Name.Local,
				path,
				attr.Value,
				actualValue,
			)
		}
	}

	return nil
}

func significantXMLChildren(node *xmlquery.Node) []*xmlquery.Node {
	children := make([]*xmlquery.Node, 0)
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == xmlquery.TextNode && strings.TrimSpace(child.Data) == "" {
			continue
		}
		children = append(children, child)
	}
	return children
}
