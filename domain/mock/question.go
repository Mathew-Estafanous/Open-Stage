package mock

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/mock"
)

// QuestionService is the mock struct for the interface implementation
// and is mainly used for unit tests.
type QuestionService struct {
	mock.Mock
}

func (m *QuestionService) FindWithId(id int) (domain.Question, error) {
	ret := m.Called(id)
	return ret.Get(0).(domain.Question), ret.Error(1)
}

func (m *QuestionService) FindAllInRoom(code string) ([]domain.Question, error) {
	ret := m.Called(code)
	return ret.Get(0).([]domain.Question), ret.Error(1)
}

func (m *QuestionStore) ChangeTotalLikes(id int, total int) error {
	ret := m.Called(id, total)
	return ret.Error(0)
}

func (m *QuestionService) Create(q *domain.Question) error {
	ret := m.Called(q)
	return ret.Error(0)
}

func (m *QuestionService) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}

// QuestionStore is the mock struct for the interface implementation
// and is mainly used for unit tests.
type QuestionStore struct {
	mock.Mock
}

func (m *QuestionStore) GetById(id int) (domain.Question, error) {
	ret := m.Called(id)
	return ret.Get(0).(domain.Question), ret.Error(1)
}

func (m *QuestionStore) GetAllInRoom(roomCode string) ([]domain.Question, error) {
	ret := m.Called(roomCode)
	return ret.Get(0).([]domain.Question), ret.Error(1)
}

func (m *QuestionStore) UpdateLikeTotal(id int, total int) error {
	ret := m.Called(id, total)
	return ret.Error(0)
}

func (m *QuestionStore) Create(q *domain.Question) error {
	ret := m.Called(q)
	return ret.Error(0)
}

func (m *QuestionStore) Delete(id int) error {
	ret := m.Called(id)
	return ret.Error(0)
}
