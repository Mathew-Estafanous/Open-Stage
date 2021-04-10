package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"net/http"
)

type roomHandler struct {
	baseHandler
	rs domain.RoomService
}

func NewRoomHandler(roomService domain.RoomService) *roomHandler {
	return &roomHandler{rs: roomService}
}

func (r roomHandler) Route(ro *mux.Router) {
	ro.HandleFunc("/room/{code}", r.getRoom).Methods("GET")
	ro.HandleFunc("/room/{code}", r.deleteRoom).Methods("DELETE")
	ro.HandleFunc("/room", r.createRoom).Methods("POST")
}

func (r roomHandler) getRoom(w http.ResponseWriter, re *http.Request) {
	roomCode := mux.Vars(re)["code"]

	room, err := r.rs.FindRoom(roomCode)
	if err != nil {
		r.error(w, err)
		return
	}

	r.respond(w, http.StatusOK, room)
}

func (r roomHandler) createRoom(w http.ResponseWriter, re *http.Request) {
	var room domain.Room
	err := json.NewDecoder(re.Body).Decode(&room)
	if err != nil {
		r.error(w, err)
		return
	}

	err = r.rs.CreateRoom(&room)
	if err != nil {
		r.error(w, err)
		return
	}

	r.respond(w, http.StatusCreated, room)
}

func (r roomHandler) deleteRoom(w http.ResponseWriter, re *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := mux.Vars(re)["code"]

	err := r.rs.DeleteRoom(code)
	if err != nil {
		r.error(w, err)
	}
}
