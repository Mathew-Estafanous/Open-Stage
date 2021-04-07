package service

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

type questionService struct {
	qStore domain.QuestionStore
}

func NewQuestionService(qStore domain.QuestionStore) domain.QuestionService {
	return &questionService{qStore}
}

func (q questionService) FindWithId(id int) (domain.Question, error) {
	question, err := q.qStore.GetById(id)
	if err != nil {
		return domain.Question{}, ErrQuestionNotFound
	}
	return question, nil
}

func (q questionService) FindAllInRoom(code string) ([]domain.Question, error) {
	//TODO: Create a check that the room code is a valid room.
	qs, err := q.qStore.GetAllInRoom(code)
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

	err := q.qStore.Create(question)
	if err != nil {
		return ErrQuestionNotCreated
	}
	return nil
}

func (q questionService) Delete(id int) error {
	err := q.qStore.Delete(id)
	if err != nil {
		return ErrInternalIssue
	}
	return nil
}

var (
	ErrQuestionNotFound     = errors.New("question with that id was not found")
	ErrQuestionMustHaveRoom = errors.New("every question must be assigned a room")
	ErrMissingQuestion      = errors.New("a question was not found")
	ErrQuestionNotCreated   = errors.New("question could not be created with given room code")
	ErrInternalIssue        = errors.New("there was an internal error")
)
