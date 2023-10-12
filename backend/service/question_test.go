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
		Host:     "Mathew",
	}

	qStore.On("GetAllInRoom", "room1").Return(foundQuestions, nil)
	rService.On("FindRoom", "room1").Return(room, nil)
	res, err := qs.FindAllInRoom("room1")
	assert.NoError(t, err)
	assert.EqualValues(t, foundQuestions, res)

	qStore.On("GetAllInRoom", "invalidRoom").Return([]domain.Question{}, errors.New("error occurred"))
	rService.On("FindRoom", "invalidRoom").Return(domain.Room{}, errRoomNotFound)
	res, err = qs.FindAllInRoom("invalidRoom")
	assert.ErrorIs(t, err, errRoomNotFound)
	assert.Nil(t, res)
}

func TestQuestionService_ChangeTotalLikes(t *testing.T) {
	qStore := new(mock.QuestionStore)
	rService := new(mock.RoomService)
	qs := NewQuestionService(qStore, rService)

	qStore.On("UpdateLikeTotal", 12, 2).Return(nil)
	qStore.On("GetById", 12).Return(domain.Question{TotalLikes: 1}, nil)
	q, err := qs.ChangeTotalLikes(12, 1)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, q.TotalLikes)

	qStore.On("GetById", 20).Return(domain.Question{}, errQuestionNotFound)
	_, err = qs.ChangeTotalLikes(20, 1)
	assert.ErrorIs(t, err, errQuestionNotFound)

	_, err = qs.ChangeTotalLikes(12, -2)
	assert.ErrorIs(t, err, errInvalidIncrement)

	qStore.On("UpdateLikeTotal", 25, 1).Return(errors.New("some sql internal error"))
	qStore.On("GetById", 25).Return(domain.Question{}, nil)
	q, err = qs.ChangeTotalLikes(25, 1)
	assert.ErrorIs(t, err, errInternalIssue)
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

	anonymousQuestion := domain.Question{
		Question:       "Is this a test?",
		AssociatedRoom: "room1",
	}
	qStore.On("Create", &domain.Question{
		Question:       anonymousQuestion.Question,
		AssociatedRoom: anonymousQuestion.AssociatedRoom,
		QuestionerName: "Anonymous",
	}).Return(nil)

	err = qs.Create(&anonymousQuestion)
	assert.NoError(t, err)
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
