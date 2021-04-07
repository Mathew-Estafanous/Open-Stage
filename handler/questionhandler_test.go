package handler

import (
	"encoding/json"
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockQuestionService struct {
	mock.Mock
}

func (m *mockQuestionService) FindWithId(id int) (domain.Question, error) {
	ret := m.Called(id)
	return ret.Get(0).(domain.Question), ret.Error(1)
}

func (m *mockQuestionService) FindAllInRoom(code string) ([]domain.Question, error) {
	ret := m.Called(code)
	return ret.Get(0).([]domain.Question), ret.Error(1)
}

func (m *mockQuestionService) Create(q *domain.Question) error {
	ret := m.Called(q)
	return ret.Error(0)
}

func (m *mockQuestionService) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func TestQuestionHandler_CreateQuestion(t *testing.T) {
	qs := new(mockQuestionService)

	question := domain.Question{
		QuestionerName: "Mathew", Question: "A question?", AssociatedRoom: "room1",
	}
	qs.On("Create", &question).Return(nil)

	j, err := json.Marshal(question)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/question", strings.NewReader(string(j)))
	assert.NoError(t, err)

	r := mux.NewRouter()
	NewQuestionHandler(qs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())

	invalidQuestion := domain.Question{QuestionerName: "Mathew", Question: "A question?"}
	j, err = json.Marshal(invalidQuestion)
	assert.NoError(t, err)

	qs.On("Create", &invalidQuestion).Return(errors.New("an error occurred"))

	req, err = http.NewRequest("POST", "/question", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}

func TestQuestionHandler_GetAllQuestionInRoom(t *testing.T) {
	qs := new(mockQuestionService)

	foundQuestions := []domain.Question{
		{
			QuestionId: 1, QuestionerName: "Mat", Question: "Is this a test?", AssociatedRoom: "room1",
		},
	}
	qs.On("FindAllInRoom", "roomCode").Return(foundQuestions, nil)

	req, err := http.NewRequest("GET", "/question/roomCode", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	NewQuestionHandler(qs).Route(r)
	r.ServeHTTP(w, req)

	j, err := json.Marshal(foundQuestions)
	assert.NoError(t, err)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())
}

func TestQuestionHandler_deleteQuestion(t *testing.T) {
	qs := new(mockQuestionService)

	qs.On("Delete", 1).Return(nil)

	req, err := http.NewRequest("DELETE", "/question/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	NewQuestionHandler(qs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	req, err = http.NewRequest("DELETE", "/question/notInt", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}
