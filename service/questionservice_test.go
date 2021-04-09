package service

import (
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockQuestionStore struct {
	mock.Mock
}

func (m *mockQuestionStore) GetById(id int) (domain.Question, error) {
	ret := m.Called(id)
	return ret.Get(0).(domain.Question), ret.Error(1)
}

func (m *mockQuestionStore) GetAllInRoom(roomCode string) ([]domain.Question, error) {
	ret := m.Called(roomCode)
	return ret.Get(0).([]domain.Question), ret.Error(1)
}

func (m *mockQuestionStore) Create(q *domain.Question) error {
	ret := m.Called(q)
	return ret.Error(0)
}

func (m *mockQuestionStore) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

func TestQuestionService_GetFromId(t *testing.T) {
	qStore := new(mockQuestionStore)
	qs := NewQuestionService(qStore)

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
	qStore := new(mockQuestionStore)
	qs := NewQuestionService(qStore)

	foundQuestions := []domain.Question{
		{
			QuestionId: 1, QuestionerName: "Mat", Question: "Is this a test?", AssociatedRoom: "room1",
		},
	}

	qStore.On("GetAllInRoom", "room1").Return(foundQuestions, nil)
	res, err := qs.FindAllInRoom("room1")
	assert.NoError(t, err)
	assert.EqualValues(t, foundQuestions, res)

	qStore.On("GetAllInRoom", "invalidRoom").Return([]domain.Question{}, errors.New("error occured"))
	res, err = qs.FindAllInRoom("invalidRoom")
	assert.ErrorIs(t, err, errInternalIssue)
	assert.Nil(t, res)
}

func TestQuestionService_Create(t *testing.T) {
	qStore := new(mockQuestionStore)
	qs := NewQuestionService(qStore)

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
	qStore := new(mockQuestionStore)
	qs := NewQuestionService(qStore)

	qStore.On("Delete", 1).Return(nil)
	err := qs.Delete(1)
	assert.NoError(t, err)

	qStore.On("Delete", 2).Return(errors.New("a database error occurred"))
	err = qs.Delete(2)
	assert.ErrorIs(t, err, errInternalIssue)
}
