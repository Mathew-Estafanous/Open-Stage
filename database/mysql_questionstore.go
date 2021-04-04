package database

import (
	"database/sql"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

type mySQLQuestionStore struct {
	db *sql.DB
}

func NewMySQLQuestionStore(db *sql.DB) domain.QuestionStore {
	return &mySQLQuestionStore{db}
}

func (m *mySQLQuestionStore) GetById(id int) (domain.Question, error) {
	row := m.db.QueryRow("SELECT question_id, question, questioner_name, total_likes, fk_room_id FROM questions WHERE question_id = ?", id)

	var question domain.Question
	err := row.Scan(&question.QuestionId, &question.Question,
		&question.QuestionerName, &question.TotalLikes, &question.AssociatedRoom)
	if err != nil {
		return domain.Question{}, err
	}
	return question, nil
}

func (m *mySQLQuestionStore) Create(q *domain.Question) error {
	r, err := m.db.Exec("INSERT INTO questions (question, questioner_name, fk_room_id) VALUES (?, ?, ?)",
		q.Question, q.QuestionerName, q.AssociatedRoom)
	if err != nil {
		return err
	}
	id, _ := r.LastInsertId()
	q.QuestionId = int(id)
	return nil
}

