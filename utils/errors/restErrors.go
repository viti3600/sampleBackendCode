package errors

import "net/http"

type RestErr struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e RestErr) AsMessage() *RestErr {
	return &RestErr{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewAuthenticationError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
