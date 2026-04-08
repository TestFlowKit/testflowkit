package queryable

import (
	"strings"

	"github.com/tidwall/gjson"
)

type JSONEngine struct {
	body []byte
}

func newJSONEngine(body []byte) Queryable {
	return &JSONEngine{body: body}
}

func (e *JSONEngine) Get(path string) (Result, error) {
	if err := validatePathSyntax(path, FormatJSON); err != nil {
		return Result{}, err
	}

	value := gjson.GetBytes(e.body, path)
	if !value.Exists() {
		return Result{}, nil
	}

	return resultFromGJSON(value), nil
}

func (e *JSONEngine) GetAll(path string) ([]Result, error) {
	if err := validatePathSyntax(path, FormatJSON); err != nil {
		return nil, err
	}

	value := gjson.GetBytes(e.body, path)
	if !value.Exists() {
		return []Result{}, nil
	}

	trimmedRaw := strings.TrimSpace(value.Raw)
	if strings.HasPrefix(trimmedRaw, "[") {
		items := value.Array()
		results := make([]Result, 0, len(items))
		for _, item := range items {
			results = append(results, resultFromGJSON(item))
		}
		return results, nil
	}

	return []Result{resultFromGJSON(value)}, nil
}

func (e *JSONEngine) Exists(path string) (bool, error) {
	if err := validatePathSyntax(path, FormatJSON); err != nil {
		return false, err
	}

	return gjson.GetBytes(e.body, path).Exists(), nil
}

func resultFromGJSON(value gjson.Result) Result {
	actualValue := value.Value()
	return Result{
		Exists: value.Exists(),
		Raw:    value.String(),
		Kind:   GetValueType(actualValue),
		Value:  actualValue,
	}
}
