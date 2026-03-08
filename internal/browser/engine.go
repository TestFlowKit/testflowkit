package browser

import "testflowkit/pkg/browser"

// Engine is the interface for browser engines.
type Engine interface {
	NewBrowser(args browser.CreationArgs) browser.Client
	Close()
}
