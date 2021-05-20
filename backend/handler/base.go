package handler

import (
	"encoding/json"
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

func newResponseError(msg string, sts int) ResponseError {
	return ResponseError{
		Msg:       msg,
		Sts:       sts,
		TimeStamp: time.Now(),
	}
}

type baseHandler struct{}

func (h baseHandler) error(w http.ResponseWriter, err error) {
	log.Print(err)
	errTypeToSts := map[domain.ErrType]int{
		domain.Internal: http.StatusInternalServerError,
		domain.NotFound: http.StatusNotFound,
		domain.Conflict: http.StatusConflict,
		domain.BadInput: http.StatusBadRequest,
	}

	var respErr ResponseError
	apiErr, ok := err.(domain.ApiError)
	if !ok {
		respErr = newResponseError("We encountered an internal error", http.StatusInternalServerError)
	} else {
		respErr = newResponseError(apiErr.Msg, errTypeToSts[apiErr.Typ])
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respErr.Sts)
	err = json.NewEncoder(w).Encode(respErr)
	if err != nil {
		log.Print(err)
	}
}

func (h baseHandler) respond(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := json.Marshal(src)
	if err != nil {
		h.error(w, err)
	}

	w.WriteHeader(code)
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
