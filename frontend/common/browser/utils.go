package browser

import (
	"cucumber/config"
	"cucumber/utils"
	"log"
)

var fc = config.FrontConfig{}

func GetElement(page Page, label string) (Element, error) {
	selectors, err := fc.GetHTMLElementSelectors(label, utils.HTMLElement)
	if err != nil {
		return nil, err
	}
	return GetElementBySelectors(page, selectors), nil
}

func GetInputElement(page Page, label string) (Element, error) {
	selectors, err := fc.GetHTMLElementSelectors(label, utils.HTMLInput)
	if err != nil {
		return nil, err
	}
	return GetElementBySelectors(page, selectors), nil
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

func GetElementByType(page Page, label string, eltType utils.ElementType) (Element, error) {
	selectors, err := fc.GetHTMLElementSelectors(label, eltType)
	if err != nil {
		return nil, err
	}

	return GetElementBySelectors(page, selectors), nil
}
