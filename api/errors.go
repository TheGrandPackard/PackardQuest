package api

type apiError struct {
	Error   string
	Message string
}

func (e *apiError) String() string {
	if e.Message != "" {
		return e.Message
	}

	return e.Error
}

func newApiError(err error) apiError {
	return apiError{Error: err.Error()}
}

func newApiErrorWithMessage(err error, msg string) apiError {
	errString := ""
	if err != nil {
		errString = err.Error()
	}

	return apiError{
		Error:   errString,
		Message: msg,
	}
}
