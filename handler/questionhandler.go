package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type QuestionHandler struct {
	qs domain.QuestionService
}

func NewQuestionHandler(qService domain.QuestionService) *QuestionHandler {
	return &QuestionHandler{qService}
}

func (q QuestionHandler) Route(router *mux.Router) {
	router.HandleFunc("/question", q.createQuestion).Methods("POST")
	router.HandleFunc("/question/{roomCode}", q.getAllQuestionsInRoom).Methods("GET")
	router.HandleFunc("/question/{questionId}", q.deleteQuestion).Methods("DELETE")
}

func (q QuestionHandler) createQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	var question domain.Question
	err := json.NewDecoder(r.Body).Decode(&question)
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

func (q QuestionHandler) getAllQuestionsInRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := mux.Vars(r)["roomCode"]

	questions, err := q.qs.GetAllWithRoomCode(code)
	encoder := json.NewEncoder(w)
	if err != nil {
		responseErr := domain.NewResponseError(
			"there was an internal error on our server", http.StatusInternalServerError)
		w.WriteHeader(responseErr.Status)
		_ = encoder.Encode(responseErr)
		return
	}

	_ = encoder.Encode(questions)
}

func (q QuestionHandler) deleteQuestion(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["questionId"]

	encoder := json.NewEncoder(w)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		responseErr := domain.NewResponseError(
			"the given question id is not a valid int", http.StatusBadRequest)
		w.WriteHeader(responseErr.Status)
		_ = encoder.Encode(responseErr)
		return
	}

	err = q.qs.Delete(idInt)
	if err != nil {
		responseErr := domain.NewResponseError(
			"there was an internal error on our server", http.StatusInternalServerError)
		w.WriteHeader(responseErr.Status)
		_ = encoder.Encode(responseErr)
	}
}
