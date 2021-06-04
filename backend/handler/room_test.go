package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestRoomHandler_GetRoom(t *testing.T) {
	rs := new(mock.RoomService)
	room := domain.Room{RoomCode: "jrHigh", Host: "Mat"}
	rs.On("FindRoom", "jrHigh").Return(room, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rooms/jrHigh", nil)
	r := mux.NewRouter()
	NewRoomHandler(rs).Route(r, r)
	r.ServeHTTP(w, req)

	roomJson, err := json.Marshal(room)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(roomJson), w.Body.String())

	rs.On("FindRoom", "wrongCode").
		Return(domain.Room{}, fmt.Errorf("%w: not found", domain.NotFound))
	req, err = http.NewRequest("GET", "/rooms/wrongCode", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}

func TestRoomHandler_CreateRoom(t *testing.T) {
	rs := new(mock.RoomService)
	room := domain.Room{RoomCode: "jrHigh", Host: "Mat", AccId: 1}
	rs.On("CreateRoom", &room).Return(nil)

	j, err := json.Marshal(room)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(string(j)))
	assert.NoError(t, err)

	r := mux.NewRouter()
	secured := r.PathPrefix("/").Subrouter()
	secured.Use(mockAuthMiddleware)
	NewRoomHandler(rs).Route(r, secured)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())

	duplicate := domain.Room{RoomCode: "duplicateCode", Host: "Mat", AccId: 1}
	rs.On("CreateRoom", &duplicate).Return(fmt.Errorf("%w: conflict", domain.Conflict))
	j, err = json.Marshal(duplicate)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/rooms", strings.NewReader(string(j)))
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusConflict, w.Code)
}

func TestRoomHandler_DeleteRoom(t *testing.T) {
	rs := new(mock.RoomService)
	rs.On("DeleteRoom", "validCode", 1).Return(nil)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/rooms/validCode", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Account", strconv.Itoa(1))

	r := mux.NewRouter()
	secured := r.PathPrefix("/").Subrouter()
	secured.Use(mockAuthMiddleware)
	NewRoomHandler(rs).Route(r, secured)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	rs.On("DeleteRoom", "wrongCode", 1).Return(fmt.Errorf("%w: not found", domain.NotFound))
	w = httptest.NewRecorder()
	req, err = http.NewRequest("DELETE", "/rooms/wrongCode", strings.NewReader(""))
	assert.NoError(t, err)
	req.Header.Set("Account", strconv.Itoa(1))

	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}

func mockAuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Account", strconv.Itoa(1))
		h.ServeHTTP(w, r)
	})
}
