package browser

import (
	"etoolse/internal/config/testsConfig"
	"log"
	"sync"
)

func GetElement(page Page, label string) (Element, error) {
	selectors, err := testsConfig.GetHTMLElementSelectors(label)
	if err != nil {
		return nil, err
	}
	return GetElementBySelectors(page, selectors), nil
}

func GetElementBySelectors(page Page, potentialSelectors []string) Element {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var foundElement Element

	ch := make(chan Element, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		wg.Add(1)

		go func() {
			defer wg.Done()

			element, err := page.GetOneBySelector(selector)
			if err != nil {
				log.Println("no element found with selector ", selector)
				return
			}
			mu.Lock()
			if foundElement == nil {
				foundElement = element
			}
			mu.Unlock()
		}()
	}

	wg.Wait()

	return foundElement
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
	potentialSelectors, _ := testsConfig.GetHTMLElementSelectors(label)
	selector := GetActiveSelector(page, potentialSelectors)
	elements, err := page.GetAllBySelector(selector)
	if err != nil {
		log.Fatal("no elements found with selector ", selector)
	}

	return len(elements)
}
