package handler

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/mock"
)

type mockQuestionService struct{
	mock.Mock
}

func (m *mockQuestionService) GetFromId(id int) (domain.Question, error) {
	panic("implement me")
}

func (m *mockQuestionService) GetAllWithRoomCode(code string) ([]domain.Question, error) {
	panic("implement me")
}

func (m *mockQuestionService) Create(q *domain.Question) error {
	panic("implement me")
}

