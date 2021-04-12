package service

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/Mathew-Estafanous/Open-Stage/domain/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuestionService_GetFromId(t *testing.T) {
	qStore := new(mock.QuestionStore)
	rService := new(mock.RoomService)
	qs := NewQuestionService(qStore, rService)

	question := domain.Question{
		QuestionId: 1, QuestionerName: "Mat", Question: "Is this a test?", AssociatedRoom: "room1",
	}

	qStore.On("GetById", 1).Return(question, nil)

	res, err := qs.FindWithId(1)
	assert.NoError(t, err)
	assert.EqualValues(t, question, res)

	qStore.On("GetById", 2).Return(domain.Question{}, errors.New("some error"))
	res, err = qs.FindWithId(2)
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestQuestionService_GetAllWithRoomCode(t *testing.T) {
	qStore := new(mock.QuestionStore)
	rService := new(mock.RoomService)
	qs := NewQuestionService(qStore, rService)

	foundQuestions := []domain.Question{
		{
			QuestionId: 1, QuestionerName: "Mat", Question: "Is this a test?", AssociatedRoom: "room1",
		},
	}

	room := domain.Room{
		RoomCode: "room1",
		Host: "Mathew",
	}

	qStore.On("GetAllInRoom", "room1").Return(foundQuestions, nil)
	rService.On("FindRoom", "room1").Return(room, nil)
	res, err := qs.FindAllInRoom("room1")
	assert.NoError(t, err)
	assert.EqualValues(t, foundQuestions, res)

	qStore.On("GetAllInRoom", "invalidRoom").Return([]domain.Question{}, errors.New("error occurred"))
	rService.On("FindRoom", "invalidRoom").Return(domain.Room{}, errRoomNotFound)
	res, err = qs.FindAllInRoom("invalidRoom")
	assert.ErrorIs(t, err, errRoomCodeNotFound)
	assert.Nil(t, res)
}

func TestQuestionService_Create(t *testing.T) {
	qStore := new(mock.QuestionStore)
	rService := new(mock.RoomService)
	qs := NewQuestionService(qStore, rService)

	validQuestion := domain.Question{
		QuestionerName: "Mat", Question: "Is this a test?", AssociatedRoom: "room1",
	}

	qStore.On("Create", &validQuestion).Return(nil)
	err := qs.Create(&validQuestion)
	assert.NoError(t, err)

	missingQuestion := domain.Question{
		QuestionerName: "Mat", AssociatedRoom: "room1",
	}
	err = qs.Create(&missingQuestion)
	assert.ErrorIs(t, err, errMissingQuestion)

	missingRoom := domain.Question{
		QuestionerName: "Mat", Question: "Is this a test?",
	}
	err = qs.Create(&missingRoom)
	assert.Error(t, err, errQuestionMustHaveRoom)
}

func TestQuestionService_Delete(t *testing.T) {
	qStore := new(mock.QuestionStore)
	rService := new(mock.RoomService)
	qs := NewQuestionService(qStore, rService)

	qStore.On("Delete", 1).Return(nil)
	err := qs.Delete(1)
	assert.NoError(t, err)

	qStore.On("Delete", 2).Return(errors.New("a database error occurred"))
	err = qs.Delete(2)
	assert.ErrorIs(t, err, errInternalIssue)
}
