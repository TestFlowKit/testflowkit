package browser

import (
	"context"
	"etoolse/internal/browser/common"
	"etoolse/internal/browser/rod"
	"etoolse/internal/config/testsconfig"
	"log"
	"sync"
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
		log.Println("no element found with selector ", sel)
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

func GetElementCount(page common.Page, label string) int {
	potentialSelectors, _ := testsconfig.GetHTMLElementSelectors(label)
	selector := getActiveSelector(page, potentialSelectors)
	elements, err := page.GetAllBySelector(selector)
	if err != nil {
		log.Fatal("no elements found with selector ", selector)
	}

	return len(elements)
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

type contextWrapper struct {
	context.Context
	cancel context.CancelFunc
}
