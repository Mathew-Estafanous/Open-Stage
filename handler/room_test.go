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

func TestJoinRoom(t *testing.T) {
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

func TestCreateRoom(t *testing.T) {
	rs := new(mockRoomService)
	room := domain.Room{RoomId: 0, RoomCode: "jrHigh", Host: "Mat"}
	rs.On("CreateRoom", &room).Return(nil)

	j, err := json.Marshal(room)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/room", strings.NewReader(string(j)))

	r := mux.NewRouter()
	NewRoomHandler(rs).Route(r)
	r.ServeHTTP(w, req)

	rj, err := json.Marshal(room)
	assert.NoError(t, err)
	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(rj), w.Body.String())

}
