package scenario

import (
	"testflowkit/internal/browser/common"
	"testflowkit/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldInstanciateCorrectlyNewFrontendContext(t *testing.T) {
	cfg := config.Config{
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
	}
	ctx := NewContext(&cfg)

	assert.InDelta(t, float64(15), ctx.frontend.timeout.Seconds(), 0.0001)
	assert.False(t, ctx.frontend.headlessMode)
	assert.InDelta(t, float64(10), ctx.frontend.slowMotion.Seconds(), 0.0001)
	assert.Nil(t, ctx.frontend.browser)
	assert.Nil(t, ctx.frontend.page)
}

func TestShouldGetPageNameByURLForInternalPage(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "http://localhost:3000",
			},
		},
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
		Frontend: config.FrontendConfig{
			Pages: map[string]string{
				"home":     "/",
				"internal": "/internal",
			},
		},
	}

	ctx := NewContext(&cfg)

	testURL := "http://localhost:3000/internal"
	ctx.frontend.page = &TestPage{
		url: testURL,
	}

	pageName, _ := ctx.getPageNameByURL(testURL)
	assert.Equal(t, "internal", pageName)
}

func TestShouldReturnErrorIfPageNameNotFound(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "https://localhost:3000",
			},
		},
	}

	ctx := NewContext(&cfg)

	pageName, err := ctx.getPageNameByURL("https://localhost:3000/unknown")
	assert.Equal(t, "page name not found", err.Error())
	assert.Equal(t, "", pageName)
}

func TestShouldGetPageNameByURLForExternalPage(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "https://localhost:3000",
			},
		},
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
		Frontend: config.FrontendConfig{
			Pages: map[string]string{
				"home":     "/",
				"internal": "/internal",
				"external": "https://testflowkit.com/external",
			},
		},
	}

	ctx := NewContext(&cfg)

	ctx.frontend.page = &TestPage{
		url: "https://testflowkit.com/external",
	}

	pageName, _ := ctx.getPageNameByURL("https://testflowkit.com/external")
	assert.Equal(t, "external", pageName)
}

func TestShouldGetPageNameByURLForInternalPageWithBaseURLContainingPath(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "https://localhost:3000",
			},
		},
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
		Frontend: config.FrontendConfig{
			Pages: map[string]string{
				"internal_with_base_url_containing_path": "https://localhost:3000/internal",
			},
		},
	}

	ctx := NewContext(&cfg)

	ctx.frontend.page = &TestPage{
		url: "https://localhost:3000/internal",
	}

	pageName, _ := ctx.getPageNameByURL("https://localhost:3000/internal")
	assert.Equal(t, "internal_with_base_url_containing_path", pageName)
}

func TestShouldGetPageNameByURLExternalWithBaseURLContainingPath(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "https://testflowkit.com/path",
			},
		},
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
		Frontend: config.FrontendConfig{
			Pages: map[string]string{
				"home":     "/",
				"internal": "/internal",
				"external": "/external",
			},
		},
	}

	ctx := NewContext(&cfg)

	ctx.frontend.page = &TestPage{
		url: "https://localhost:3000/path",
	}

	pageName, _ := ctx.getPageNameByURL("https://testflowkit.com/path/external")
	assert.Equal(t, "external", pageName)
}

func TestShouldGetPageNameByURLInternalWithBaseURLContainingPath(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "https://localhost:3000/path",
			},
		},
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
		Frontend: config.FrontendConfig{
			Pages: map[string]string{
				"home":     "/",
				"internal": "/internal",
			},
		},
	}

	ctx := NewContext(&cfg)

	ctx.frontend.page = &TestPage{
		url: "https://localhost:3000/path/internal",
	}

	pageName, _ := ctx.getPageNameByURL("https://localhost:3000/path/internal")
	assert.Equal(t, "internal", pageName)
}

func TestShouldGetPageNameByURLWithVariableForInternalPage(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "https://localhost:3000",
			},
		},
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
		Frontend: config.FrontendConfig{
			Pages: map[string]string{
				"home":    "/",
				"details": "/product/:id",
			},
		},
	}

	ctx := NewContext(&cfg)

	ctx.frontend.page = &TestPage{
		url: "https://localhost:3000/product/abc",
	}

	pageName, _ := ctx.getPageNameByURL("https://localhost:3000/product/abc")
	assert.Equal(t, "details", pageName)
}

func TestShouldGetPageNameByVariableSupportBaseURLDifferentFromPageURL(t *testing.T) {
	cfg := config.Config{
		ActiveEnvironment: "local",
		Environments: map[string]config.Environment{
			"local": {
				FrontendBaseURL: "https://localhost:3000",
			},
		},
		Settings: config.GlobalSettings{
			DefaultTimeout: 15000,
			Headless:       false,
			SlowMotion:     10000,
		},
		Frontend: config.FrontendConfig{
			Pages: map[string]string{
				"home":             "/",
				"sentence-details": "sentences/:id",
			},
		},
	}

	ctx := NewContext(&cfg)

	ctx.frontend.page = &TestPage{
		url: "https://testflowkit.com/sentences/abc",
	}

	pageName, _ := ctx.getPageNameByURL("https://testflowkit.com/sentences/abc")
	assert.NotEqual(t, "sentence-details", pageName)
}

type TestPage struct {
	url string
	common.Page
}

func (p *TestPage) GetInfo() common.PageInfo {
	return common.PageInfo{
		URL: p.url,
	}
}
