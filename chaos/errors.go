package chaos

import "log"

const retriableErrorString = "retriableError"
const notFoundErrorString = "resourceNotFound"

func shouldRetryError(err error) bool {
	shouldRetry := err.Error() == retriableErrorString
	log.Printf("[WARN] shouldRetryError, shouldRetry: %v", shouldRetry)
	return shouldRetry
}

func shouldIgnoreError(err error) bool {
	return err.Error() == notFoundErrorString
}
