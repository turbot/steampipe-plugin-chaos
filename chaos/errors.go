package chaos

var retriableError = "retriableError"
var notFoundError = "resourceNotFound"

func shouldRetryError(err error) bool {
	return err.Error() == retriableError
}

func shouldIgnoreError(err error) bool {
	return err.Error() == notFoundError
}
