package scenario

import (
	"context"
	"testflowkit/internal/browser/common"
	"testflowkit/internal/config"
)

func GetPage(ctx context.Context) (common.Page, string, error) {
	scenarioCtx := MustFromContext(ctx)
	return scenarioCtx.GetCurrentPage()
}

func GetPageOnly(ctx context.Context) (common.Page, error) {
	scenarioCtx := MustFromContext(ctx)
	return scenarioCtx.GetCurrentPageOnly()
}

func GetConfig(ctx context.Context) *config.Config {
	scenarioCtx := MustFromContext(ctx)
	return scenarioCtx.GetConfig()
}

func OpenPage(ctx context.Context, url string) {
	scenarioCtx := MustFromContext(ctx)
	scenarioCtx.OpenNewPage(url)
}

func SetPage(ctx context.Context, page common.Page) error {
	scenarioCtx := MustFromContext(ctx)
	return scenarioCtx.SetCurrentPage(page)
}

func GetKeyboard(ctx context.Context) common.Keyboard {
	scenarioCtx := MustFromContext(ctx)
	return scenarioCtx.GetCurrentPageKeyboard()
}

func GetPages(ctx context.Context) []common.Page {
	scenarioCtx := MustFromContext(ctx)
	return scenarioCtx.GetPages()
}

func UpdatePageName(ctx context.Context) {
	scenarioCtx := MustFromContext(ctx)
	scenarioCtx.UpdatePageNameIfNeeded()
}
