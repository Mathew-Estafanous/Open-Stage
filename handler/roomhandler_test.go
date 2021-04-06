package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/service"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockRoomService struct {
	mock.Mock
}

func (m *mockRoomService) CreateRoom(room *domain.Room) error {
	ret := m.Called(room)
	return ret.Error(0)
}

func (m *mockRoomService) FindRoom(roomCode string) (domain.Room, error) {
	ret := m.Called(roomCode)
	return ret.Get(0).(domain.Room), ret.Error(1)
}

func (m *mockRoomService) DeleteRoom(code string) error {
	ret := m.Called(code)
	return ret.Error(0)
}

func TestRoomHandler_GetRoom(t *testing.T) {
	rs := new(mockRoomService)
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

	rs.On("FindRoom", "wrongCode").Return(domain.Room{}, service.ErrRoomNotFound)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/room/wrongCode", nil)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}

func TestRoomHandler_CreateRoom(t *testing.T) {
	rs := new(mockRoomService)
	room := domain.Room{RoomCode: "jrHigh", Host: "Mat"}
	rs.On("CreateRoom", &room).Return(nil)

	j, err := json.Marshal(room)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/room", strings.NewReader(string(j)))

	r := mux.NewRouter()
	NewRoomHandler(rs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())

	duplicate := domain.Room{RoomCode: "duplicateCode", Host: "Mat"}
	rs.On("CreateRoom", &duplicate).Return(service.ErrDuplicateRoom)
	j, err = json.Marshal(duplicate)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/room", strings.NewReader(string(j)))
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusConflict, w.Code)
}

func TestRoomHandler_DeleteRoom(t *testing.T) {
	rs := new(mockRoomService)
	rs.On("DeleteRoom", "validCode").Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/room/validCode", strings.NewReader(""))
	r := mux.NewRouter()
	NewRoomHandler(rs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	rs.On("DeleteRoom", "wrongCode").Return(service.ErrRoomNotDeleted)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/room/wrongCode", strings.NewReader(""))
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusNotFound, w.Code)
}
