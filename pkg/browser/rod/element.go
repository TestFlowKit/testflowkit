package rod

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"testflowkit/pkg/browser"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type rodElement struct {
	element *rod.Element
}

func (e *rodElement) Input(text string) error {
	err := e.element.Input(text)
	if err != nil {
		return err
	}
	return nil
}

func (e *rodElement) Click() error {
	return e.element.Click(
		proto.InputMouseButtonLeft,
		1,
	)
}

func (e *rodElement) DoubleClick() error {
	const numberOfClicks = 2
	return e.element.Click(proto.InputMouseButtonLeft, numberOfClicks)
}

func (e *rodElement) RightClick() error {
	return e.element.Click(proto.InputMouseButtonRight, 1)
}

func (e *rodElement) TextContent() string {
	return e.element.MustText()
}

func (e *rodElement) InputValue() string {
	return e.element.MustText()
}

func (e *rodElement) IsVisible() bool {
	visible, err := e.element.Visible()
	if err != nil {
		return false
	}
	return visible
}

func (e *rodElement) SelectByText(options []string) error {
	return e.element.Select(options, true, rod.SelectorTypeRegex)
}

func (e *rodElement) SelectByValue(values []string) error {
	var dropdownsTexts []string
	optionsElts, getEltserr := e.element.Elements("option")
	if getEltserr != nil {
		return getEltserr
	}

	for _, option := range optionsElts {
		attrValue, getAttrErr := option.Attribute("value")
		if getAttrErr != nil {
			continue
		}

		isExpected := slices.Contains(values, *attrValue)
		if !isExpected {
			continue
		}
		dropdownsTexts = append(dropdownsTexts, option.MustText())
	}

	if len(dropdownsTexts) == 0 {
		return errors.New("no option found")
	}

	if len(dropdownsTexts) < len(values) {
		return fmt.Errorf("only Options with values %v found", dropdownsTexts)
	}

	return e.SelectByText(dropdownsTexts)
}

func (e *rodElement) SelectByIndex(optionIndex int) error {
	optionsElts, err := e.element.Elements("option")
	if err != nil {
		return err
	}

	if optionIndex < 0 || optionIndex >= len(optionsElts) {
		return errors.New("invalid option index, max index is " + strconv.Itoa(len(optionsElts)-1))
	}

	return e.SelectByText([]string{optionsElts[optionIndex].MustText()})
}

func (e *rodElement) GetAttributeValue(property string, kind reflect.Kind) any {
	value := e.element.MustProperty(property)

	if kind == reflect.Bool {
		return value.Bool()
	}

	if kind == reflect.String {
		return value.String()
	}

	return nil
}

func (e *rodElement) Hover() error {
	return e.element.Hover()
}

func (e *rodElement) IsChecked() bool {
	value, err := e.element.Property("checked")
	if err != nil {
		return false
	}
	return value.Bool()
}

func (e *rodElement) Clear() error {
	return e.element.Input("")
}

func (e *rodElement) UploadFile(filePath string) error {
	if err := e.checkFilesExist([]string{filePath}); err != nil {
		return err
	}
	return e.element.SetFiles([]string{filePath})
}

func (e *rodElement) UploadMultipleFiles(filePaths []string) error {
	if err := e.checkFilesExist(filePaths); err != nil {
		return err
	}
	return e.element.SetFiles(filePaths)
}

func (e *rodElement) checkFilesExist(filePaths []string) error {
	if len(filePaths) == 0 {
		return errors.New("no file paths provided")
	}

	notFoundFiles := []string{}
	for _, filePath := range filePaths {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			notFoundFiles = append(notFoundFiles, filePath)
		}
	}

	if len(notFoundFiles) > 0 {
		return fmt.Errorf("files do not exist: %v", notFoundFiles)
	}

	return nil
}

func (e *rodElement) ScrollIntoView() error {
	return e.element.ScrollIntoView()
}

func (e *rodElement) Check() error {
	return e.Click()
}

func (e *rodElement) Uncheck() error {
	if !e.IsChecked() {
		return nil
	}
	return e.Click()
}

func newRodElement(element *rod.Element) browser.Element {
	return &rodElement{element: element}
}
