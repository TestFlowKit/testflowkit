package browser

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testflowkit/internal/config"
	"testflowkit/pkg/browser"
	"testflowkit/pkg/logger"
)

func GetElementByLabel(page browser.Page, pageName, label string) (browser.Element, error) {
	cfg, err := config.Get()

	if err != nil {
		return nil, err
	}

	selectors := cfg.GetElementSelectors(pageName, label)
	if len(selectors) == 0 {
		return nil, errors.New("no selectors found for element " + label)
	}

	elt := getElementBySelectors(page, selectors)
	if elt == nil {
		return nil, errors.New("element not found")
	}

	return elt, nil
}

func getElementBySelectors(page browser.Page, potentialSelectors []config.Selector) browser.Element {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan browser.Element, 1)
	defer close(ch)

	var mu sync.RWMutex
	for _, selector := range potentialSelectors {
		go searchForSelector(searchSelectorParams{
			ctx:      contextWrapper{Context: ctx, cancel: cancel},
			mu:       &mu,
			page:     page,
			selector: selector,
			resultCh: ch,
		})
	}

	<-ctx.Done()

	cancel()
	return <-ch
}

// searchSelectorParams holds the parameters for searching a selector.
type searchSelectorParams struct {
	ctx      contextWrapper
	mu       *sync.RWMutex
	page     browser.Page
	selector config.Selector
	resultCh chan<- browser.Element
}

func searchForSelector(params searchSelectorParams) {
	var elt browser.Element
	var err error

	value := params.selector.String()
	if params.selector.IsXPath() {
		elt, err = params.page.GetOneByXPath(value)
	} else {
		elt, err = params.page.GetOneBySelector(value)
	}

	if err != nil {
		logger.Warn(fmt.Sprintf("element not found with %s selector %s", params.selector.Type, value), []string{
			"Please fix the selector in the configuration file",
			"Please verify that page is accessible",
		})

		select {
		case <-params.ctx.Done():
			return
		default:
			params.resultCh <- nil
			params.ctx.cancel()
			return
		}
	}

	if elt != nil {
		params.mu.Lock()
		defer params.mu.Unlock()

		select {
		case <-params.ctx.Done():
			return
		default:
			params.resultCh <- elt
			params.ctx.cancel()
		}
	}
}

func getActiveSelector(page browser.Page, potentialSelectors []config.Selector) config.Selector {
	ch := make(chan config.Selector, 1)
	defer close(ch)

	for _, selector := range potentialSelectors {
		go func(sel config.Selector) {
			var exists bool
			if sel.IsXPath() {
				// For XPath, we need to check if the element exists
				elt, err := page.GetOneByXPath(sel.Value)
				exists = err == nil && elt != nil
			} else {
				exists = page.HasSelector(sel.Value)
			}

			if exists {
				ch <- sel
			}
		}(selector)
	}

	return <-ch
}

func GetElementCount(page browser.Page, pageName, label string) int {
	cfg, err := config.Get()
	if err != nil {
		return 0
	}

	potentialSelectors := cfg.GetElementSelectors(pageName, label)
	selector := getActiveSelector(page, potentialSelectors)

	var elements []browser.Element
	var err2 error

	if selector.IsXPath() {
		elt, getByXpathErr := page.GetOneByXPath(selector.Value)
		if getByXpathErr == nil {
			elements = []browser.Element{elt}
		}
	} else {
		elements, err2 = page.GetAllBySelector(selector.Value)
	}

	if err2 != nil {
		msg := fmt.Sprintf("Error getting elements with %s selector %s", selector.Type, selector.Value)
		logger.Error(msg, []string{
			"Incorrect selector defined in the configuration file",
		}, []string{"Check the selector in the configuration file"})
	}

	return len(elements)
}

type contextWrapper struct {
	context.Context
	cancel context.CancelFunc
}
