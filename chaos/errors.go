package chaos

const retriableError = "retriableError"
const notFoundError = "resourceNotFound"

func shouldRetryError(err error) bool {
	return err.Error() == retriableError
}

func shouldIgnoreError(err error) bool {
	return err.Error() == notFoundError
}
