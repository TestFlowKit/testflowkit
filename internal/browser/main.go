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

func CreateInstance(headlessMode bool, timeout, slowMotion time.Duration, incognitoMode bool) common.Browser {
	return rod.New(headlessMode, timeout, slowMotion, incognitoMode)
}

func GetElementByLabel(page common.Page, pageName, label string) (common.Element, error) {
	cfg, err := config.Get()

	if err != nil {
		return nil, err
	}

	selectors := cfg.GetElementSelectors(pageName, config.GetLabelKey(label))
	if len(selectors) == 0 {
		return nil, errors.New("no selectors found for element " + label)
	}

	element := getElementBySelectors(page, selectors)
	if element == nil {
		return nil, errors.New("element not found")
	}

	return element, nil
}

func getElementBySelectors(page common.Page, potentialSelectors []string) common.Element {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan common.Element, 1)
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

func searchForSelector(ctx contextWrapper, mu *sync.RWMutex, p common.Page, sel string, ch chan<- common.Element) {
	element, err := p.GetOneBySelector(sel)
	if err != nil {
		logger.Warn(fmt.Sprintf("element not found with selector %s", sel), []string{
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

	if element != nil {
		mu.Lock()
		defer mu.Unlock()

		select {
		case <-ctx.Done():
			return
		default:
			ch <- element
			ctx.cancel()
		}
	}
}

func getActiveSelector(page common.Page, potentialSelectors []string) string {
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

func GetElementCount(page common.Page, pageName, label string) int {
	cfg, err := config.Get()
	if err != nil {
		return 0
	}

	potentialSelectors := cfg.GetElementSelectors(pageName, label)
	selector := getActiveSelector(page, potentialSelectors)
	elements, err := page.GetAllBySelector(selector)
	if err != nil {
		msg := fmt.Sprintf("Error getting elements with selector %s", selector)
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
