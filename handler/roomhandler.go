package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	"github.com/gorilla/mux"
	"net/http"
)

type RoomHandler struct {
	rs domain.RoomService
}

func NewRoomHandler(roomService domain.RoomService) *RoomHandler {
	return &RoomHandler{rs: roomService}
}

func (r RoomHandler) Route(router *mux.Router) {
	router.HandleFunc("/room/{code}", r.getRoom).Methods("GET")
	router.HandleFunc("/room", r.createRoom).Methods("POST")
	router.HandleFunc("/room/{code}", r.deleteRoom).Methods("DELETE")
}

func (r RoomHandler) getRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomCode := mux.Vars(re)["code"]

	room, err := r.rs.FindRoom(roomCode)
	if err != nil {
		responseError := domain.NewResponseError(err.Error(), http.StatusNotFound)
		w.WriteHeader(responseError.Status)
		_ = json.NewEncoder(w).Encode(responseError)
		return
	}
	_ = json.NewEncoder(w).Encode(room)
}

func (r RoomHandler) createRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var room domain.Room
	encoder := json.NewEncoder(w)
	err := json.NewDecoder(re.Body).Decode(&room)
	if err != nil {
		responseErr := domain.NewResponseError(
			"Request body not formatted properly", http.StatusBadRequest)
		w.WriteHeader(responseErr.Status)
		_ = encoder.Encode(responseErr)
		return
	}

	err = r.rs.CreateRoom(&room)
	if err != nil {
		status := http.StatusConflict
		if err == service.ErrHostNotAssigned {
			status = http.StatusBadRequest
		}

		responseErr := domain.NewResponseError(
			err.Error(), status)
		w.WriteHeader(responseErr.Status)
		_ = encoder.Encode(responseErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = encoder.Encode(room)
}

func (r RoomHandler) deleteRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := mux.Vars(re)["code"]
	err := r.rs.DeleteRoom(code)
	if err != nil {
		responseErr := domain.NewResponseError(err.Error(), http.StatusNotFound)
		w.WriteHeader(responseErr.Status)
		_ = json.NewEncoder(w).Encode(responseErr)
		return
	}
}
