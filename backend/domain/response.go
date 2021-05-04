package domain

import (
	"net/http"
	"time"
)

type ResponseError struct {
	// The server error message.
	Msg string `json:"message"`
	// The https status error.
	Sts int `json:"status"`
	// Time in which the error has occurred.
	TimeStamp time.Time `json:"timestamp"`
}

func (r ResponseError) Error() string {
	return r.Msg
}

func NewResponseError(msg string, sts int) ResponseError {
	return ResponseError{
		Msg:       msg,
		Sts:       sts,
		TimeStamp: time.Now(),
	}
}

func NotFound(msg string) ResponseError {
	return NewResponseError(msg, http.StatusNotFound)
}

func InternalServerError(msg string) ResponseError {
	if msg == "" {
		msg = "We encountered an internal error while processing the request."
	}
	return NewResponseError(msg, http.StatusInternalServerError)
}

func BadRequest(msg string) ResponseError {
	if msg == "" {
		msg = "The given request was in a bad format."
	}
	return NewResponseError(msg, http.StatusBadRequest)
}

func Conflict(msg string) ResponseError {
	if msg == "" {
		msg = "There was a conflict while processing the request."
	}
	return NewResponseError(msg, http.StatusConflict)
}
