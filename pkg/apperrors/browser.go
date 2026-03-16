package apperrors

import "errors"

// ErrElementNotFound is returned when a browser element cannot be located.
var ErrElementNotFound = errors.New("element not found")

// ErrNoOptionFound is returned when no matching option exists in a select element.
var ErrNoOptionFound = errors.New("no option found")

// ErrNoFilePaths is returned when a file upload step receives no file paths.
var ErrNoFilePaths = errors.New("no file paths provided")

// ErrNoCurrentPage is returned when a frontend step requires an open page but none is active.
var ErrNoCurrentPage = errors.New("no current page available")

// ErrNoNewWindow is returned when a step waits for a new browser window that never appears.
var ErrNoNewWindow = errors.New("no new window detected")
