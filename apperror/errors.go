package apperror

import (
	"log"
	"net/http"
)

type ErrType int

const (
	ErrBadRequest ErrType = iota + 1
	ErrForbidden
	ErrNotFound
	ErrInternalServer
	ErrTimeout
	ErrUnauthorized
)

type CustomError struct {
	ErrType    ErrType
	LogMessage string
	err        error
}

func (cu *CustomError) Error() string {
	return cu.err.Error()
}

func NewCustomError(errType ErrType, logMessage string, err error) *CustomError {
	log.Println(logMessage)
	return &CustomError{
		ErrType:    errType,
		LogMessage: logMessage,
		err:        err,
	}
}

func (cu *CustomError) ErrCode() int {
	switch cu.ErrType {
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrForbidden:
		return http.StatusForbidden
	case ErrNotFound:
		return http.StatusNotFound
	case ErrTimeout:
		return http.StatusGatewayTimeout
	case ErrUnauthorized:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
