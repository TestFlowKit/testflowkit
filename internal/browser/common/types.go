package common

import (
	"reflect"

	"github.com/go-rod/rod/lib/input"
)

type Browser interface {
	NewPage(url string) Page
	GetPages() []Page
	Close()
}

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
}

type PageInfo struct {
	URL string
}

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
	GetPropertyValue(property string, kind reflect.Kind) any
}

type Keyboard interface {
	Press(key input.Key) error
}
