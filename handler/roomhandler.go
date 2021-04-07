package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type roomHandler struct {
	rs domain.RoomService
}

func NewRoomHandler(roomService domain.RoomService) *roomHandler {
	return &roomHandler{rs: roomService}
}

func (r roomHandler) Route(router *mux.Router) {
	router.HandleFunc("/room/{code}", r.getRoom).Methods("GET")
	router.HandleFunc("/room/{code}", r.deleteRoom).Methods("DELETE")
	router.HandleFunc("/room", r.createRoom).Methods("POST")
}

func (r roomHandler) getRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomCode := mux.Vars(re)["code"]

	room, err := r.rs.FindRoom(roomCode)
	if err != nil {
		responseErr := domain.NewResponseError(err.Error(), http.StatusNotFound)
		writeResponseErr(w, responseErr)
		return
	}

	if err = json.NewEncoder(w).Encode(room); err != nil {
		log.Print(err)
	}
}

func (r roomHandler) createRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var room domain.Room
	err := json.NewDecoder(re.Body).Decode(&room)
	if err != nil {
		responseErr := domain.NewResponseError(
			"Request body not formatted properly", http.StatusBadRequest)
		writeResponseErr(w, responseErr)
		return
	}

	err = r.rs.CreateRoom(&room)
	if err != nil {
		status := http.StatusConflict
		if err == service.ErrHostNotAssigned {
			status = http.StatusBadRequest
		}

		responseErr := domain.NewResponseError(err.Error(), status)
		writeResponseErr(w, responseErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(room); err != nil {
		log.Print(err)
	}
}

func (r roomHandler) deleteRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := mux.Vars(re)["code"]

	err := r.rs.DeleteRoom(code)
	if err != nil {
		responseErr := domain.NewResponseError(err.Error(), http.StatusNotFound)
		writeResponseErr(w, responseErr)
	}
}
