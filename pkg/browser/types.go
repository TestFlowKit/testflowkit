package browser

import (
	"reflect"
	"time"
)

// Key represents an abstract keyboard key.
type Key string

// Keyboard key constants.
const (
	KeyEnter      Key = "Enter"
	KeyTab        Key = "Tab"
	KeyDelete     Key = "Delete"
	KeyEscape     Key = "Escape"
	KeySpace      Key = "Space"
	KeyArrowUp    Key = "Arrow Up"
	KeyArrowRight Key = "Arrow Right"
	KeyArrowDown  Key = "Arrow Down"
	KeyArrowLeft  Key = "Arrow Left"
)

// Client represents a browser client instance.
type Client interface {
	NewPage(url string) Page
	GetPages() []Page
	Close()
}

// Page represents a browser page.
type Page interface {
	GetOneBySelector(selector string) (Element, error)
	GetAllBySelector(selector string) ([]Element, error)
	GetOneByXPath(xpath string) (Element, error)
	GetOneByTextContent(text string) (Element, error)
	WaitLoading()
	Refresh()
	GetInfo() PageInfo
	Focus()
	Back()
	Screenshot() ([]byte, error)
	GetKeyboard() Keyboard
	HasSelector(selector string) bool
	ExecuteJS(js string, args ...any) string
	SetTimeout(timeout time.Duration)
}

// PageInfo contains information about a browser page.
type PageInfo struct {
	URL   string
	Title string
}

// Element represents a DOM element.
type Element interface {
	Click() error
	DoubleClick() error
	RightClick() error
	Hover() error
	Input(text string) error
	Clear() error
	ScrollIntoView() error
	IsChecked() bool
	SelectByText([]string) error
	SelectByValue([]string) error
	SelectByIndex(int) error
	IsVisible() bool
	TextContent() string
	GetAttributeValue(attribute string, kind reflect.Kind) any
	UploadFile(filePath string) error
	UploadMultipleFiles(filePaths []string) error
}

// Keyboard represents a keyboard input handler.
type Keyboard interface {
	Press(key Key) error
}
