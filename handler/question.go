package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type questionHandler struct {
	baseHandler
	qs domain.QuestionService
}

func NewQuestionHandler(qService domain.QuestionService) *questionHandler {
	return &questionHandler{qs: qService}
}

func (q questionHandler) Route(r *mux.Router) {
	r.HandleFunc("/question", q.createQuestion).Methods("POST")
	r.HandleFunc("/question/{roomCode}", q.getAllQuestionsInRoom).Methods("GET")
	r.HandleFunc("/question/{questionId}", q.deleteQuestion).Methods("DELETE")
}

func (q questionHandler) createQuestion(w http.ResponseWriter, r *http.Request) {
	var question domain.Question
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		q.error(w, err)
		return
	}

	err = q.qs.Create(&question)
	if err != nil {
		q.error(w, err)
		return
	}

	q.respond(w, http.StatusCreated, question)
}

func (q questionHandler) getAllQuestionsInRoom(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["roomCode"]

	questions, err := q.qs.FindAllInRoom(code)
	if err != nil {
		q.error(w, err)
		return
	}

	q.respond(w, http.StatusOK, questions)
}

func (q questionHandler) deleteQuestion(w http.ResponseWriter, r *http.Request)  {
	id := mux.Vars(r)["questionId"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = domain.BadRequest("Given question Id is not a valid int type.")
		q.error(w, err)
		return
	}

	err = q.qs.Delete(idInt)
	if err != nil {
		q.error(w, err)
	}
}
