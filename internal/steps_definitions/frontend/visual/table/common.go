package table

import (
	"errors"
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

	trs, _ := page.GetAllBySelector("tr")

	idx := slices.IndexFunc(trs, func(elt common.Element) bool {
		textContent := elt.TextContent()
		for _, value := range content {
			if !strings.Contains(textContent, value) {
				return false
			}
		}
		return true
	})

	if idx == -1 {
		return nil, errors.New("row not found with the following values: " + strings.Join(content, ", "))
	}

	element := trs[idx]
	if !element.IsVisible() {
		return nil, errors.New("row exists but it is not visible")
	}

	return element, nil
}
