package service

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

type questionService struct {
	qStore   domain.QuestionStore
	rService domain.RoomService
}

func NewQuestionService(qStore domain.QuestionStore, rService domain.RoomService) domain.QuestionService {
	return &questionService{qStore, rService}
}

func (q questionService) FindWithId(id int) (domain.Question, error) {
	question, err := q.qStore.GetById(id)
	if err != nil {
		return domain.Question{}, errQuestionNotFound
	}
	return question, nil
}

func (q questionService) FindAllInRoom(code string) ([]domain.Question, error) {
	_, err := q.rService.FindRoom(code)
	if err != nil {
		return nil, errRoomCodeNotFound
	}

	qs, err := q.qStore.GetAllInRoom(code)
	if err != nil {
		return nil, errInternalIssue
	}
	return qs, nil
}

func (q questionService) Create(question *domain.Question) error {
	if question.AssociatedRoom == "" {
		return errQuestionMustHaveRoom
	}

	if question.Question == "" {
		return errMissingQuestion
	}

	if question.QuestionerName == "" {
		question.QuestionerName = "Anonymous"
	}

	err := q.qStore.Create(question)
	if err != nil {
		return errQuestionNotCreated
	}
	return nil
}

func (q questionService) Delete(id int) error {
	err := q.qStore.Delete(id)
	if err != nil {
		return errInternalIssue
	}
	return nil
}

var (
	errRoomCodeNotFound     = domain.NotFound("The given room code was not found.")
	errQuestionNotFound     = domain.NotFound("A question with that id was not found.")
	errQuestionMustHaveRoom = domain.BadRequest("Every question must be assigned a room.")
	errMissingQuestion      = domain.BadRequest("A question was not provided.")
	errQuestionNotCreated   = domain.BadRequest("The question could not be created using the room code.")
	errInternalIssue        = domain.InternalServerError("")
)
