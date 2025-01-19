package browser

import (
	"etoolse/config"
	"log"
)

var fc = config.FrontConfig{}

func GetElement(page Page, label string) (Element, error) {
	selectors, err := fc.GetHTMLElementSelectors(label)
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
	potentialSelectors, _ := fc.GetHTMLElementSelectors(label)
	selector := GetActiveSelector(page, potentialSelectors)
	elements, err := page.GetAllBySelector(selector)
	if err != nil {
		log.Fatal("no elements found with selector ", selector)
	}

	return len(elements)
}
