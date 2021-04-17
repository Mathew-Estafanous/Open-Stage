package postgres

import (
	"database/sql"
	"errors"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
)

type mySQLQuestionStore struct {
	db *sql.DB
}

func NewQuestionStore(db *sql.DB) domain.QuestionStore {
	return &mySQLQuestionStore{db}
}

func (m *mySQLQuestionStore) GetById(id int) (domain.Question, error) {
	row := m.db.QueryRow("SELECT question_id, question, questioner_name, total_likes, fk_room_code FROM questions WHERE question_id = $1", id)

	var question domain.Question
	err := row.Scan(&question.QuestionId, &question.Question,
		&question.QuestionerName, &question.TotalLikes, &question.AssociatedRoom)
	if err != nil {
		return domain.Question{}, err
	}
	return question, nil
}

func (m mySQLQuestionStore) GetAllInRoom(code string) ([]domain.Question, error) {
	rows, err := m.db.Query("SELECT question_id, question, questioner_name, total_likes, fk_room_code FROM questions WHERE fk_room_code = $1", code)
	if err != nil {
		return nil, err
	}

	question := make([]domain.Question, 0)
	for rows.Next() {
		var q domain.Question
		err := rows.Scan(&q.QuestionId, &q.Question,
			&q.QuestionerName, &q.TotalLikes, &q.AssociatedRoom)
		if err != nil {
			return nil, err
		}
		question = append(question, q)
	}
	return question, nil
}

func (m *mySQLQuestionStore) UpdateLikeTotal(id int, total int) error {
	r, err := m.db.Exec("UPDATE questions SET total_likes = $1 WHERE question_id = $2", total, id)
	if err != nil {
		return err
	}

	a, err := r.RowsAffected()
	if err != nil {
		return err
	}

	if a == 0 {
		return errors.New("no rows were altered as expected")
	}
	return nil
}

func (m *mySQLQuestionStore) Create(q *domain.Question) error {
	r, err := m.db.Exec("INSERT INTO questions (question, questioner_name, fk_room_code) VALUES ($1, $2, $3)",
		q.Question, q.QuestionerName, q.AssociatedRoom)
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return err
	}
	q.QuestionId = int(id)
	return nil
}

func (m *mySQLQuestionStore) Delete(id int) error {
	_, err := m.db.Exec("DELETE FROM questions WHERE question_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
