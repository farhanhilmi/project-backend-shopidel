package util

import (
	"fmt"
)

type ErrorCode int

const (
	BadRequest ErrorCode = iota
	NotFound
	Unauthorized
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrWrongCredentials  = &CustomError{Message: "wrong email or password", Code: Unauthorized}
	ErrNoRecordFound     = &CustomError{Message: "not found", Code: NotFound}
	ErrBadPIN            = &CustomError{Message: "length minimum length is 6 and should be numeric", Code: BadRequest}
	ErrInvalidInput      = &CustomError{Message: "invalid input", Code: BadRequest}
	ErrWalletAlreadySet  = &CustomError{Message: "wallet already activated", Code: BadRequest}
	ErrWalletNotSet      = &CustomError{Message: "your wallet has not been activated", Code: BadRequest}
	ErrWalletPINNotMatch = &CustomError{Message: "your wallet pin doesnt match", Code: BadRequest}
	ErrInvalidPassword   = &CustomError{Message: "invalid password", Code: Unauthorized}
	ErrSameWalletPIN     = &CustomError{Message: "your new wallet PIN must not be the same as your current PIN"}
)
