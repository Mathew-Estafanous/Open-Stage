package service

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

type questionService struct {
	questionStore domain.QuestionStore
}

func NewQuestionService(qStore domain.QuestionStore) domain.QuestionService {
	return &questionService{qStore}
}

func (q questionService) GetFromId(id int) (domain.Question, error) {
	question, err := q.questionStore.GetById(id)
	if err != nil {
		return domain.Question{}, ErrQuestionNotFound
	}
	return question, nil
}

func (q questionService) GetAllWithRoomCode(code string) ([]domain.Question, error) {
	//TODO: Create a check that the room code is a valid room.
	qs, err := q.questionStore.GetAllForRoom(code)
	if err != nil {
		return nil, ErrInternalIssue
	}
	return qs, nil
}

func (q questionService) Create(question *domain.Question) error {
	if question.AssociatedRoom == "" {
		return ErrQuestionMustHaveRoom
	}

	if question.Question == "" {
		return ErrMissingQuestion
	}

	if question.QuestionerName == "" {
		question.QuestionerName = "Anonymous"
	}

	err := q.questionStore.Create(question)
	if err != nil {
		return ErrQuestionNotCreated
	}
	return nil
}

var (
	ErrQuestionNotFound     = errors.New("question with that id was not found")
	ErrQuestionMustHaveRoom = errors.New("every question must be assigned a room")
	ErrMissingQuestion      = errors.New("a question was not found")
	ErrQuestionNotCreated   = errors.New("the question was not created")
	ErrInternalIssue        = errors.New("there was an internal error")
)
