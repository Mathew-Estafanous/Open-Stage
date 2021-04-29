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

type updateLike struct {
	Id         int `json:"question_id"`
	TotalLikes int `json:"total_likes"`
}

func NewQuestionHandler(qService domain.QuestionService) *questionHandler {
	return &questionHandler{qs: qService}
}

func (q questionHandler) Route(r *mux.Router) {
	r.HandleFunc("/questions", q.createQuestion).Methods("POST")
	r.HandleFunc("/questions", q.updateTotalLikes).Methods("PUT")
	r.HandleFunc("/questions/{roomCode}", q.getAllQuestionsInRoom).Methods("GET")
	r.HandleFunc("/questions/{questionId}", q.deleteQuestion).Methods("DELETE")
}

// swagger:route POST /questions Questions createQuestion
//
// Create new question in associated room.
//
// Uploads a new question to the room. The questioner's name is optional
// and will be left as "Anonymous" by default.
//
// Responses:
//   201: questionResponse
//   400: errorResponse
//   500: errorResponse
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

func (q questionHandler) updateTotalLikes(w http.ResponseWriter, r *http.Request) {
	var body updateLike
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		q.error(w, err)
		return
	}

	err = q.qs.ChangeTotalLikes(body.Id, body.TotalLikes)
	if err != nil {
		q.error(w, err)
		return
	}
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

func (q questionHandler) deleteQuestion(w http.ResponseWriter, r *http.Request) {
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
