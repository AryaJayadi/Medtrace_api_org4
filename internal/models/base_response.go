package models

import "fmt"

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseValueResponse[T any] struct {
	Success bool       `json:"success"`
	Value   *T         `json:"value,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
}

type BaseListResponse[T any] struct {
	Success bool       `json:"success"`
	List    []*T       `json:"list"`
	Error   *ErrorInfo `json:"error,omitempty"`
}

func SuccessValueResponse[T any](value T) BaseValueResponse[T] {
	return BaseValueResponse[T]{
		Success: true,
		Value:   &value,
		Error:   nil,
	}
}

func ErrorValueResponse[T any](code int, format string, args ...any) BaseValueResponse[T] {
	return BaseValueResponse[T]{
		Success: false,
		Value:   nil,
		Error: &ErrorInfo{
			Code:    code,
			Message: fmt.Sprintf(format, args...),
		},
	}
}

func SuccessListResponse[T any](list []*T) BaseListResponse[T] {
	if list == nil {
		list = make([]*T, 0)
	}
	return BaseListResponse[T]{
		Success: true,
		List:    list,
		Error:   nil,
	}
}

func ErrorListResponse[T any](code int, format string, args ...any) BaseListResponse[T] {
	return BaseListResponse[T]{
		Success: false,
		List:    nil,
		Error: &ErrorInfo{
			Code:    code,
			Message: fmt.Sprintf(format, args...),
		},
	}
}
