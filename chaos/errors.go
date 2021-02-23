package chaos

var retriableError = "retriableError"
var notFoundError = "ResourceNotFound"

func shouldRetryError(err error) bool {
	return err.Error() == retriableError
}

func shouldIgnoreError(err error) bool {
	return err.Error() == notFoundError
}
