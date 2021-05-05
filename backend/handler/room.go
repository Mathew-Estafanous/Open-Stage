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
	ro.HandleFunc("/rooms/{code}", r.getRoom).Methods("GET")
	ro.HandleFunc("/rooms/{code}", r.deleteRoom).Methods("DELETE")
	ro.HandleFunc("/rooms", r.createRoom).Methods("POST")
}

// swagger:route GET /rooms/{code} Rooms getCode
//
// Get room by code.
//
// Simply fetches the room that equals the code that was passed in.
//
// Responses:
//  200: roomResponse
//  404: errorResponse
//  500: errorResponse
func (r roomHandler) getRoom(w http.ResponseWriter, re *http.Request) {
	roomCode := mux.Vars(re)["code"]

	room, err := r.rs.FindRoom(roomCode)
	if err != nil {
		r.error(w, err)
		return
	}

	r.respond(w, http.StatusOK, room)
}

// swagger:route POST /rooms Rooms createRoom
//
// Create a new room.
//
// Startup a new room with an assigned host. However, the room code is
// not required and if left empty will be randomly generated. If a code
// is already in use by another room, a 409 Conflict will occur.
//
// Responses:
//  201: roomResponse
//  400: errorResponse
//  409: errorResponse
//  500: errorResponse
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

// swagger:route DELETE /rooms/{code} Rooms delCode
//
// Delete a room by code.
//
// Responses:
//  200: description: OK
//  500: errorResponse
func (r roomHandler) deleteRoom(w http.ResponseWriter, re *http.Request) {
	code := mux.Vars(re)["code"]

	err := r.rs.DeleteRoom(code)
	if err != nil {
		r.error(w, err)
	}
}