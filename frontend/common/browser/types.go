package browser

import "reflect"

type Browser interface {
	NewPage(url string) Page
}

type Page interface {
	GetOneBySelector(selector string) (Element, error)
	GetAllBySelector(selector string) ([]Element, error)
	GetOneByXPath(xpath string) (Element, error)
	WaitLoading()
	GetInfo() PageInfo
	GetKeyboard() Keyboard
	HasSelector(selector string) bool
	ExecuteJS(js string, args ...any) string
}

type PageInfo struct {
	URL string
}

type Element interface {
	Click() error
	Input(text string) error
	Select([]string) error
	IsVisible() bool
	TextContent() string
	GetPropertyValue(property string, kind reflect.Kind) any
}

type Keyboard interface {
	PressEnter() error
}
