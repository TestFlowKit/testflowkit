package playwright

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
	"testflowkit/pkg/browser"

	pw "github.com/playwright-community/playwright-go"
)

type playwrightElement struct {
	locator pw.Locator
}

func (e *playwrightElement) Input(text string) error {
	return e.locator.Fill(text)
}

func (e *playwrightElement) Click() error {
	return e.locator.Click()
}

func (e *playwrightElement) DoubleClick() error {
	return e.locator.Dblclick()
}

func (e *playwrightElement) RightClick() error {
	return e.locator.Click(pw.LocatorClickOptions{
		Button: pw.MouseButtonRight,
	})
}

func (e *playwrightElement) TextContent() string {
	textContent, err := e.locator.TextContent()
	if err != nil {
		return ""
	}
	return textContent
}

func (e *playwrightElement) Check() error {
	return e.locator.Check()
}

func (e *playwrightElement) Uncheck() error {
	return e.locator.Uncheck()
}

func (e *playwrightElement) InputValue() string {
	textContent, err := e.locator.InputValue()
	if err != nil {
		return ""
	}
	return textContent
}

func (e *playwrightElement) IsVisible() bool {
	visible, err := e.locator.IsVisible()
	if err != nil {
		return false
	}
	return visible
}

func (e *playwrightElement) SelectByText(options []string) error {
	for _, option := range options {
		_, err := e.locator.SelectOption(pw.SelectOptionValues{
			Labels: &[]string{option},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *playwrightElement) SelectByValue(values []string) error {
	// Get all option elements
	optionsLocator := e.locator.Locator("option")
	count, err := optionsLocator.Count()
	if err != nil {
		return err
	}

	var dropdownsTexts []string
	for i := range count {
		optionLocator := optionsLocator.Nth(i)
		attrValue, errGetAttr := optionLocator.GetAttribute("value")
		if errGetAttr != nil {
			continue
		}

		if attrValue == "" {
			continue
		}

		isExpected := slices.Contains(values, attrValue)
		if !isExpected {
			continue
		}

		textContent, errGetAttr := optionLocator.TextContent()
		if errGetAttr != nil {
			continue
		}

		if textContent != "" {
			dropdownsTexts = append(dropdownsTexts, textContent)
		}
	}

	if len(dropdownsTexts) == 0 {
		return errors.New("no option found")
	}

	if len(dropdownsTexts) < len(values) {
		return fmt.Errorf("only Options with values %v found", dropdownsTexts)
	}

	return e.SelectByText(dropdownsTexts)
}

func (e *playwrightElement) SelectByIndex(optionIndex int) error {
	optionsLocator := e.locator.Locator("option")
	count, err := optionsLocator.Count()
	if err != nil {
		return err
	}

	if optionIndex < 0 || optionIndex >= count {
		return errors.New("invalid option index, max index is " + strconv.Itoa(count-1))
	}

	optionLocator := optionsLocator.Nth(optionIndex)
	textContent, err := optionLocator.TextContent()
	if err != nil {
		return err
	}

	if textContent == "" {
		return errors.New("option text not found")
	}

	return e.SelectByText([]string{textContent})
}

func (e *playwrightElement) GetAttributeValue(property string, kind reflect.Kind) any {
	attrValue, err := e.locator.GetAttribute(property)
	if err != nil {
		return nil
	}

	if attrValue == "" {
		return nil
	}

	if kind == reflect.Bool {
		return attrValue == "true"
	}

	if kind == reflect.String {
		return attrValue
	}

	return nil
}

func (e *playwrightElement) Hover() error {
	return e.locator.Hover()
}

func (e *playwrightElement) IsChecked() bool {
	checked, err := e.locator.IsChecked()
	if err != nil {
		return false
	}
	return checked
}

func (e *playwrightElement) Clear() error {
	return e.locator.Fill("")
}

func (e *playwrightElement) UploadFile(filePath string) error {
	if err := e.checkFilesExist([]string{filePath}); err != nil {
		return err
	}
	return e.locator.SetInputFiles(filePath)
}

func (e *playwrightElement) UploadMultipleFiles(filePaths []string) error {
	if err := e.checkFilesExist(filePaths); err != nil {
		return err
	}

	err := e.locator.SetInputFiles(filePaths)
	if err != nil {
		return err
	}

	return nil
}

func (e *playwrightElement) checkFilesExist(filePaths []string) error {
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

func (e *playwrightElement) ScrollIntoView() error {
	return e.locator.ScrollIntoViewIfNeeded()
}

func newPlaywrightElement(locator pw.Locator) browser.Element {
	return &playwrightElement{locator: locator}
}
