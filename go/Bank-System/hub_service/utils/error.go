package utils

import (
	"encoding/json"
	"errors"
	"log"
)

type Error struct {
	Code    ErrorCode `json:"code"`
	Err     error     `json:"err"`
	Message string    `json:"message"`
}

func (e *Error) Error() string {
	json, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
	}
	return string(json)
}

func Wrap(err error, info string) *Error {
	return &Error{
		Message: info,
		Err:     err,
	}
}

func NewError() error {
	return &Error{
		Code:    2,
		Err:     errors.New("Unknown"),
		Message: "Unknown error",
	}
}
