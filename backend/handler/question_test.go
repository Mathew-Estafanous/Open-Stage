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

func TestQuestionHandler_createQuestion(t *testing.T) {
	qs := new(mock.QuestionService)

	question := domain.Question{
		QuestionerName: "Mathew", Question: "A question?", AssociatedRoom: "room1",
	}
	qs.On("Create", &question).Return(nil)

	j, err := json.Marshal(question)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/questions", strings.NewReader(string(j)))
	assert.NoError(t, err)

	r := mux.NewRouter()
	NewQuestionHandler(qs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())

	invalidQuestion := domain.Question{QuestionerName: "Mathew", Question: "A question?"}
	j, err = json.Marshal(invalidQuestion)
	assert.NoError(t, err)

	qs.On("Create", &invalidQuestion).Return(domain.ApiError{Msg: "", Typ: domain.BadInput})

	req, err = http.NewRequest("POST", "/questions", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}

func TestQuestionHandler_getAllQuestionInRoom(t *testing.T) {
	qs := new(mock.QuestionService)

	foundQuestions := []domain.Question{
		{
			QuestionId: 1, QuestionerName: "Mat", Question: "Is this a test?", AssociatedRoom: "room1",
		},
	}
	qs.On("FindAllInRoom", "roomCode").Return(foundQuestions, nil)

	req, err := http.NewRequest("GET", "/questions/roomCode", nil)
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

func TestQuestionHandler_updateTotalLikes(t *testing.T) {
	qs := new(mock.QuestionService)

	updatedQuestion := domain.Question{
		QuestionId: 12,
		Question:   "How is everyone doing?",
		TotalLikes: 1,
	}

	qs.On("ChangeTotalLikes", 12, 1).Return(updatedQuestion, nil)

	validBody := UpdateLike{
		Id:            12,
		LikeIncrement: 1,
	}
	j, err := json.Marshal(validBody)
	assert.NoError(t, err)

	req, err := http.NewRequest("PUT", "/questions", strings.NewReader(string(j)))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	NewQuestionHandler(qs).Route(r)
	r.ServeHTTP(w, req)

	j, err = json.Marshal(updatedQuestion)
	assert.NoError(t, err)

	assert.EqualValues(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(j), w.Body.String())
}

func TestQuestionHandler_deleteQuestion(t *testing.T) {
	qs := new(mock.QuestionService)

	qs.On("Delete", 1).Return(nil)

	req, err := http.NewRequest("DELETE", "/questions/1", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	NewQuestionHandler(qs).Route(r)
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusOK, w.Code)

	req, err = http.NewRequest("DELETE", "/questions/notInt", nil)
	assert.NoError(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.EqualValues(t, http.StatusBadRequest, w.Code)
}
