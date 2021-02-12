package chaos

import (
	"github.com/turbot/go-kit/helpers"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

var chaosError = "retriableError"

// function which returns an ErrorPredicate for AWS API calls
func shouldRetryError(shouldRetryErrors []string) plugin.ErrorPredicate {
	return func(err error) bool {
		return helpers.StringSliceContains(shouldRetryErrors, chaosError)
	}
}
