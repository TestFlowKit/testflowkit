package testsconfig

import "fmt"

var elements = map[string][]string{
	"Document 1": {".file-item:nth-child(1)"},
	"Document 2": {".file-item:nth-child(2)"},
}

func GetHTMLElementSelectors(element string) ([]string, error) {
	selectors, exists := elements[element]
	if !exists {
		return nil, fmt.Errorf("element %s not found", element)
	}
	return selectors, nil
}

func IsElementDefined(element string) bool {
	_, exists := elements[element]
	return exists
} 