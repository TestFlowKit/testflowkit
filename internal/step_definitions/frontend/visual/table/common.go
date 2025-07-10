package table

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
	"testflowkit/internal/browser/common"
)

func getTableRowByCellsContent(currentPage common.Page, cellsContent []string) (common.Element, error) {
	return getTableRowOrHeaderByCellsContent(currentPage, "td", cellsContent)
}

func getTableHeaderByCellsContent(currentPage common.Page, cellsContent []string) (common.Element, error) {
	return getTableRowOrHeaderByCellsContent(currentPage, "th", cellsContent)
}

func getTableRowOrHeaderByCellsContent(page common.Page, selector string, content []string) (common.Element, error) {
	allowedValues := []string{"th", "td"}
	if !slices.Contains(allowedValues, selector) {
		log.Panicf("only %s allowed", strings.Join(allowedValues, ", "))
	}

	var xpathParts []string
	for _, value := range content {
		xpathParts = append(xpathParts, fmt.Sprintf("%s[contains(text(), '%s')]", selector, value))
	}

	xPath := fmt.Sprintf("//tr[%s]", strings.Join(xpathParts, " and "))

	element, err := page.GetOneByXPath(xPath)
	if err != nil {
		return nil, err
	}

	if element == nil {
		return nil, errors.New("row not found")
	}

	if !element.IsVisible() {
		// TODO: better message
		return nil, errors.New("row is not visible")
	}

	return element, nil
}
