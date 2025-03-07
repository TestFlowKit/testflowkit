package browser

import (
	"context"
	"fmt"
	"sync"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/browser/rod"
	"testflowkit/internal/config/testsconfig"
	"testflowkit/pkg/logger"
	"time"
)

func CreateInstance(headlessMode bool, timeout, slowMotion time.Duration, incognitoMode bool) common.Browser {
	return rod.New(headlessMode, timeout, slowMotion, incognitoMode)
}

func GetElementByLabel(page common.Page, label string) (common.Element, error) {
	selectors, err := testsconfig.GetHTMLElementSelectors(label)
	if err != nil {
		return nil, err
	}
	return getElementBySelectors(page, selectors), nil
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

func GetElementCount(page common.Page, label string) int {
	potentialSelectors, _ := testsconfig.GetHTMLElementSelectors(label)
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
