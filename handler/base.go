package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"log"
	"net/http"
)

type baseHandler struct{}

func (h baseHandler) error(w http.ResponseWriter, err error) {
	log.Print(err)
	respErr, ok := err.(domain.ResponseError)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		w.WriteHeader(http.StatusInternalServerError)
		msg := domain.InternalServerError("")
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			log.Println(err)
		}
	}

	w.WriteHeader(code)
	_, err = w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
