package handler

import (
	"encoding/json"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type questionHandler struct {
	qs domain.QuestionService
}

func NewQuestionHandler(qService domain.QuestionService) *questionHandler {
	return &questionHandler{qService}
}

func (q questionHandler) Route(router *mux.Router) {
	router.HandleFunc("/question", q.createQuestion).Methods("POST")
	router.HandleFunc("/question/{roomCode}", q.getAllQuestionsInRoom).Methods("GET")
	router.HandleFunc("/question/{questionId}", q.deleteQuestion).Methods("DELETE")
}

func (q questionHandler) createQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)

	var question domain.Question
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		handleError(w, err)
		return
	}

	err = q.qs.Create(&question)
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = e.Encode(question); err != nil {
		log.Print(err)
	}
}

func (q questionHandler) getAllQuestionsInRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	code := mux.Vars(r)["roomCode"]

	questions, err := q.qs.FindAllInRoom(code)
	e := json.NewEncoder(w)
	if err != nil {
		handleError(w, err)
		return
	}

	if err = e.Encode(questions); err != nil {
		log.Print(err)
	}
}

func (q questionHandler) deleteQuestion(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["questionId"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = domain.BadRequest("Given question Id is not a valid int type.")
		handleError(w, err)
		return
	}

	err = q.qs.Delete(idInt)
	if err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	log.Print(err)
	respErr, ok := err.(domain.ResponseError)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(respErr.Sts)
	err = json.NewEncoder(w).Encode(respErr)
	if err != nil {
		log.Print(err)
	}
}
