package browser

import (
	"cucumber/config"
	"log"
)

var fc = config.FrontConfig{}

func GetElement(page Page, label string) Element {
	selectors, _ := fc.GetElementSelectors(label)
	return GetElementBySelectors(page, selectors)
}

func GetInputElement(page Page, label string) Element {
	selectors, err := fc.GetInputSelectors(label)
	if err != nil {
		return nil
	}
	return GetElementBySelectors(page, selectors)
}

func GetElementBySelectors(page Page, potentialSelectors []string) Element {
	ch := make(chan Element, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		go func() {
			element, _ := page.GetOneBySelector(selector)
			ch <- element
		}()
	}

	return <-ch
}

func GetActiveSelector(page Page, potentialSelectors []string) string {
	ch := make(chan string, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		go func() {
			exists := page.HasSelector(selector)
			if exists {
				ch <- selector
			}
		}()
	}

	return <-ch
}

func GetElementCount(page Page, label string) int {
	potentialSelectors, _ := fc.GetElementSelectors(label)
	selector := GetActiveSelector(page, potentialSelectors)
	elements, err := page.GetAllBySelector(selector)
	if err != nil {
		log.Fatal("no elements found with selector ", selector)
	}

	return len(elements)
}
