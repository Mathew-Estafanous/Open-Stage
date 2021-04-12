package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRoomHandler_GetRoom(t *testing.T) {
	rs := new(mock.RoomService)
	room := domain.Room{RoomId: 1, RoomCode: "jrHigh", Host: "Mat"}
	rs.On("FindRoom", "jrHigh").Return(room, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/room/jrHigh", nil)
	r := mux.NewRouter()
	NewRoomHandler(rs).Route(r)
	r.ServeHTTP(w, req)

	roomJson, err := json.Marshal(room)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(roomJson), w.Body.String())

	rs.On("FindRoom", "wrongCode").
		Return(domain.Room{}, domain.NotFound(""))
	req, err = http.NewRequest("GET", "/room/wrongCode", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}

func TestRoomHandler_CreateRoom(t *testing.T) {
	rs := new(mock.RoomService)
	room := domain.Room{RoomCode: "jrHigh", Host: "Mat"}
	rs.On("CreateRoom", &room).Return(nil)

	j, err := json.Marshal(room)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/room", strings.NewReader(string(j)))
	assert.NoError(t, err)

	r := mux.NewRouter()
	NewRoomHandler(rs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())

	duplicate := domain.Room{RoomCode: "duplicateCode", Host: "Mat"}
	rs.On("CreateRoom", &duplicate).Return(domain.Conflict(""))
	j, err = json.Marshal(duplicate)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/room", strings.NewReader(string(j)))
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusConflict, w.Code)
}

func TestRoomHandler_DeleteRoom(t *testing.T) {
	rs := new(mock.RoomService)
	rs.On("DeleteRoom", "validCode").Return(nil)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/room/validCode", strings.NewReader(""))
	assert.NoError(t, err)

	r := mux.NewRouter()
	NewRoomHandler(rs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	rs.On("DeleteRoom", "wrongCode").Return(domain.NotFound(""))
	w = httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", "/room/wrongCode", strings.NewReader(""))
	assert.NoError(t, err)

	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}
