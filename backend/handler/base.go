package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/handle_err"
	"log"
	"net/http"
)

type baseHandler struct{}

func (h baseHandler) error(w http.ResponseWriter, err error) {
	respError := handle_err.ToHttp(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respError.Sts)
	err = json.NewEncoder(w).Encode(respError)
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
