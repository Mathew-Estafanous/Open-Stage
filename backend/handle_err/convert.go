package handle_err

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"log"
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

var errTypeToSts = map[domain.Code]int{
	domain.Internal:     http.StatusInternalServerError,
	domain.NotFound:     http.StatusNotFound,
	domain.Conflict:     http.StatusConflict,
	domain.BadInput:     http.StatusBadRequest,
	domain.Unauthorized: http.StatusUnauthorized,
	domain.Forbidden:    http.StatusForbidden,
}

func ToHttp(err error) ResponseError {
	log.Println(err)
	var code domain.Code
	if errors.As(err, &code) {
		return ResponseError{Msg: err.Error(), Sts: errTypeToSts[code], TimeStamp: time.Now()}
	}
	return ResponseError{Msg: "we encountered an internal error", Sts: http.StatusInternalServerError, TimeStamp: time.Now()}
}
