package main

import (
	"errors"
	"fmt"
)

type HttpError struct {
	Code    int
	Desc    string
	Message string
}

func NewHttpError(code int, desc, message string) *HttpError {
	return &HttpError{
		Code:    code,
		Desc:    desc,
		Message: message,
	}
}

func (he *HttpError) Error() string {
	return fmt.Sprintf("errcode:%s, errdesc:%s, detail message:%s", he.Code, he.Desc, he.Message)
}

func main() {
	errors.New("err")
}