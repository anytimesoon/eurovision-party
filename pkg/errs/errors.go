package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewInvalidError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

type commonMessages struct {
	BadlyFormedObject string
	DBFail            string
	NotFound          string
	NotDeleted        string
	NotUpdated        string
	NotCreated        string
	Unauthorized      string
}

var Common commonMessages = commonMessages{
	BadlyFormedObject: "Failed to read request. Please make sure your object is correctly formed.",
	DBFail:            "Something went wrong when accessing the database",
	NotFound:          "Could not find ",
	NotDeleted:        "Something went wrong when deleting ",
	NotUpdated:        "Something went wrong when updating ",
	NotCreated:        "Something went wrong when creating ",
	Unauthorized:      "You are not allowed to do that!",
}
