package browser

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/browser/rod"
	"testflowkit/internal/config"
	"testflowkit/pkg/logger"
	"time"
)

type page = common.Page
type element = common.Element
type Browser = common.Browser

func CreateInstance(headlessMode bool, thinkTime time.Duration, incognitoMode bool) Browser {
	return rod.New(headlessMode, thinkTime, incognitoMode)
}

func GetElementByLabel(page page, pageName, label string) (element, error) {
	cfg, err := config.Get()

	if err != nil {
		return nil, err
	}

	selectors := cfg.GetElementSelectors(pageName, config.GetLabelKey(label))
	if len(selectors) == 0 {
		return nil, errors.New("no selectors found for element " + label)
	}

	elt := getElementBySelectors(page, selectors)
	if elt == nil {
		return nil, errors.New("element not found")
	}

	return elt, nil
}

func getElementBySelectors(page page, potentialSelectors []config.Selector) element {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan element, 1)
	defer close(ch)

	var mu sync.RWMutex
	for _, selector := range potentialSelectors {
		go searchForSelector(contextWrapper{
			Context: ctx,
			cancel:  cancel,
		}, &mu, page, selector, ch)
	}

	<-ctx.Done()

	cancel()
	return <-ch
}

func searchForSelector(ctx contextWrapper, mu *sync.RWMutex, p page, selector config.Selector, ch chan<- element) {
	var elt element
	var err error

	value := selector.String()
	if selector.IsXPath() {
		elt, err = p.GetOneByXPath(value)
	} else {
		elt, err = p.GetOneBySelector(value)
	}

	if err != nil {
		logger.Warn(fmt.Sprintf("element not found with %s selector %s", selector.Type, value), []string{
			"Please fix the selector in the configuration file",
			"Please verify that page is accessible",
		})

		select {
		case <-ctx.Done():
			return
		default:
			ch <- nil
			ctx.cancel()
			return
		}
	}

	if elt != nil {
		mu.Lock()
		defer mu.Unlock()

		select {
		case <-ctx.Done():
			return
		default:
			ch <- elt
			ctx.cancel()
		}
	}
}

func getActiveSelector(page page, potentialSelectors []config.Selector) config.Selector {
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

func GetElementCount(page page, pageName, label string) int {
	cfg, err := config.Get()
	if err != nil {
		return 0
	}

	potentialSelectors := cfg.GetElementSelectors(pageName, label)
	selector := getActiveSelector(page, potentialSelectors)

	var elements []element
	var err2 error

	if selector.IsXPath() {
		elt, getByXpathErr := page.GetOneByXPath(selector.Value)
		if getByXpathErr == nil {
			elements = []element{elt}
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
