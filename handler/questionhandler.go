package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"net/http"
)

type QuestionHandler struct {
	qs domain.QuestionService
}

func NewQuestionHandler(qService domain.QuestionService) *QuestionHandler {
	return &QuestionHandler{qService}
}

func (q QuestionHandler) Route(router *mux.Router) {
	router.HandleFunc("/question", q.CreateQuestion).Methods("POST")
	router.HandleFunc("/question/{roomCode}", q.GetAllQuestionInRoom).Methods("GET")
}

func (q QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var question domain.Question
	err := json.NewDecoder(r.Body).Decode(&question)

	encoder := json.NewEncoder(w)
	if err != nil {
		responseErr := domain.NewResponseError(
			"There was an issue parsing the request body.", http.StatusInternalServerError)
		w.WriteHeader(responseErr.Status)
		_ = encoder.Encode(responseErr)
		return
	}

	err = q.qs.Create(&question)
	if err != nil {
		responseErr := domain.NewResponseError(err.Error(), http.StatusBadRequest)
		w.WriteHeader(responseErr.Status)
		_ = encoder.Encode(responseErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = encoder.Encode(question)
}

func (q QuestionHandler) GetAllQuestionInRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := mux.Vars(r)["roomCode"]

	questions, err := q.qs.GetAllWithRoomCode(code)
	encoder := json.NewEncoder(w)
	if err != nil {
		responseErr := domain.NewResponseError(
			"there was an internal error on our server", http.StatusInternalServerError)
		_ = encoder.Encode(responseErr)
		return
	}

	_ = encoder.Encode(questions)
}