package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// UpdateLike represents the question's new like total.
//
// swagger:model updateLikes
type UpdateLike struct {
	// The ID of the question
	//
	// required: true
	// example: 3452
	Id int `json:"question_id"`

	// Whether
	//
	// required: true
	// example: -1
	LikeIncrement int `json:"like_increment"`
}

// NewQuestion represents the request body of new questions.
//
// swagger:model newQuestion
type NewQuestion struct {
	// The question that was asked.
	//
	// required: true
	// example: Do you like this API?
	Question string `json:"question"`

	// The room the question was asked in.
	//
	// required: true
	// example: conference20
	Room string `json:"associated_room"`

	// Name of the questioner.
	//
	// default: Anonymous
	// example: Mathew
	Questioner string `json:"questioner_name"`
}

type QuestionHandler struct {
	qs domain.QuestionService
}

func NewQuestionHandler(qService domain.QuestionService) *QuestionHandler {
	return &QuestionHandler{qs: qService}
}

func (q QuestionHandler) Route(r *mux.Router) {
	r.HandleFunc("/questions", q.createQuestion).Methods("POST")
	r.HandleFunc("/questions", q.updateTotalLikes).Methods("PUT")
	r.HandleFunc("/questions/{roomCode}", q.getAllQuestionsInRoom).Methods("GET")
	r.HandleFunc("/questions/{questionId}", q.deleteQuestion).Methods("DELETE")
}

// swagger:route POST /questions Questions createQuestion
//
// Create new question in room.
//
// Uploads a new question to the room. The questioner's name is optional
// and will be left as "Anonymous" by default.
//
// Responses:
//
//	201: questionResponse
//	400: errorResponse
//	500: errorResponse
func (q QuestionHandler) createQuestion(w http.ResponseWriter, r *http.Request) {
	var body *NewQuestion
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, err)
		return
	}

	question := &domain.Question{
		Question:       body.Question,
		QuestionerName: body.Questioner,
		AssociatedRoom: body.Room,
	}

	err = q.qs.Create(question)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithCode(w, http.StatusCreated, question)
}

// swagger:route PUT /questions Questions updateLikes
//
// Update question's total number of likes.
//
// Updates the total # of likes for the question with the matching question_id.
//
// Responses:
//
//	200: questionResponse
//	404: errorResponse
//	400: errorResponse
//	500: errorResponse
func (q QuestionHandler) updateTotalLikes(w http.ResponseWriter, r *http.Request) {
	var body UpdateLike
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respondWithError(w, err)
		return
	}

	question, err := q.qs.ChangeTotalLikes(body.Id, body.LikeIncrement)
	if err != nil {
		respondWithError(w, err)
		return
	}
	respondWithCode(w, http.StatusOK, question)
}

// swagger:route GET /questions/{roomCode} Questions roomCode
//
// Get all questions in a room.
//
// Uses the given room code and retrieves all questions that have been posted.
//
// Responses:
//
//	200: multiQuestionResponse
//	404: errorResponse
//	500: errorResponse
func (q QuestionHandler) getAllQuestionsInRoom(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["roomCode"]

	questions, err := q.qs.FindAllInRoom(code)
	if err != nil {
		respondWithError(w, err)
		return
	}

	respondWithCode(w, http.StatusOK, questions)
}

// swagger:route DELETE /questions/{questionId} Questions questionId
//
// # Delete a question by ID
//
// Uses the given question ID to delete the question with that ID.
//
// Responses:
//
//	200: description: OK - Question has been properly deleted.
//	500: errorResponse
func (q QuestionHandler) deleteQuestion(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["questionId"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("%w: given question id is not a valid int type", domain.BadInput)
		respondWithError(w, err)
		return
	}

	err = q.qs.Delete(idInt)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithCode(w, http.StatusOK, &GenericResponse{
			Msg:  "successfully deleted",
			Code: http.StatusOK,
		})
	}
}
