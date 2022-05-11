package chaos

import (
	"context"
	"log"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

const retriableErrorString = "retriableError"
const notFoundErrorString = "resourceNotFound"

func shouldRetryErrorLegacy(err error) bool {
	return shouldRetryError(nil, nil, nil, err)
}

func shouldIgnoreErrorLegacy(err error) bool {
	return shouldIgnoreError(nil, nil, nil, err)
}

func shouldRetryError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	shouldRetry := err.Error() == retriableErrorString
	log.Printf("[INFO] shouldRetryError, shouldRetry: %v, err: %s", shouldRetry, err.Error())
	return shouldRetry
}

func shouldIgnoreError(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
	shouldIgnore := err.Error() == notFoundErrorString
	log.Printf("[INFO] shouldIgnoreError, shouldIgnore: %v, err: %s", shouldIgnore, err.Error())
	return shouldIgnore
}
