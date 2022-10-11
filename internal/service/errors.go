package service

type ValidationError struct {
	msg       string
	errorsMap map[string]string
}

func NewValidationError(message string, errorsMap map[string]string) error {
	return ValidationError{
		msg:       message,
		errorsMap: errorsMap,
	}
}

func (v ValidationError) Error() string {
	return v.msg
}

func (v ValidationError) ErrorsMap() map[string]string {
	return v.errorsMap
}

type AuthorizationError struct {
	message string
}

func NewAuthorizationError(message string) error {
	return AuthorizationError{
		message: message,
	}
}

func (e AuthorizationError) Error() string {
	return e.message
}
