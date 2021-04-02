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
	router.HandleFunc("/room/{code}", r.JoinRoom).Methods("GET")
	router.HandleFunc("/room", r.CreateRoom).Methods("POST")
	router.HandleFunc("/room/{code}", r.DeleteRoom).Methods("DELETE")
}

func (r *RoomHandler) JoinRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roomCode := mux.Vars(re)["code"]

	room, err := r.rs.FindRoom(roomCode)
	if err != nil {
		responseError := domain.NewResponseError(err.Error(), http.StatusNotFound)
		w.WriteHeader(responseError.Status)
		json.NewEncoder(w).Encode(responseError)
		return
	}
	json.NewEncoder(w).Encode(room)
}

func (r RoomHandler) CreateRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var room domain.Room
	encoder := json.NewEncoder(w)
	err := json.NewDecoder(re.Body).Decode(&room)
	if err != nil {
		responseErr := domain.NewResponseError(
			"Request body not formatted properly", http.StatusBadRequest)
		w.WriteHeader(responseErr.Status)
		encoder.Encode(responseErr)
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
		encoder.Encode(responseErr)
		return
	}
	w.WriteHeader(http.StatusCreated)
	encoder.Encode(room)
}

func (r *RoomHandler) DeleteRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := mux.Vars(re)["code"]
	err := r.rs.DeleteRoom(code)
	if err != nil {
		responseErr := domain.NewResponseError(err.Error(), http.StatusNotFound)
		w.WriteHeader(responseErr.Status)
		json.NewEncoder(w).Encode(responseErr)
		return
	}
}
