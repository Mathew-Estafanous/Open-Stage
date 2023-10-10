package handler

import (
	"encoding/json"
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"log"
	"net/http"
	"time"
)

func respondWithError(w http.ResponseWriter, err error) {
	respError := ToHttp(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respError.Sts)
	err = json.NewEncoder(w).Encode(respError)
	if err != nil {
		log.Print(err)
	}
}

func respondWithCode(w http.ResponseWriter, code int, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := json.Marshal(src)
	if err != nil {
		respondWithError(w, err)
	}

	w.WriteHeader(code)
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

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
