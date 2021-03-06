package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresQuestionStore_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	mQuestion := domain.Question{
		QuestionId: 2, Question: "How you doing?", QuestionerName: "Mathew", TotalLikes: 0, AssociatedRoom: "room1",
	}

	row := sqlmock.NewRows([]string{"question_id", "question", "questioner_name", "total_likes", "fk_room_code"}).
		AddRow(mQuestion.QuestionId, mQuestion.Question, mQuestion.QuestionerName,
			mQuestion.TotalLikes, mQuestion.AssociatedRoom)

	questionId := 2
	query := `SELECT question_id, question, questioner_name, total_likes, fk_room_code
				FROM questions WHERE question_id = ?`
	mock.ExpectQuery(query).WithArgs(questionId).WillReturnRows(row)

	qStore := NewQuestionStore(db)
	q, err := qStore.GetById(2)

	assert.NoError(t, err)
	assert.EqualValues(t, mQuestion, q)
}

func TestPostgresQuestionStore_GetAllInRoom(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	qs := []domain.Question{
		{
			QuestionId: 1, Question: "How you doing?", QuestionerName: "Mathew",
			AssociatedRoom: "room1", TotalLikes: 1,
		},
		{
			QuestionId: 2, Question: "Is everything good?", QuestionerName: "Anonymous",
			AssociatedRoom: "room1", TotalLikes: 3,
		},
	}

	rows := sqlmock.NewRows([]string{"question_id", "question", "questioner_name", "total_likes", "fk_room_code"}).
		AddRow(qs[0].QuestionId, qs[0].Question, qs[0].QuestionerName, qs[0].TotalLikes, qs[0].AssociatedRoom).
		AddRow(qs[1].QuestionId, qs[1].Question, qs[1].QuestionerName, qs[1].TotalLikes, qs[1].AssociatedRoom)

	query := `SELECT question_id, question, questioner_name, total_likes, fk_room_code
				FROM questions WHERE fk_room_code = ?`
	mock.ExpectQuery(query).WithArgs("room1").WillReturnRows(rows)

	qStore := NewQuestionStore(db)
	result, err := qStore.GetAllInRoom(qs[0].AssociatedRoom)
	assert.NoError(t, err)
	assert.EqualValues(t, qs, result)
}

func TestPostgresQuestionStore_UpdateLikeTotal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("there was an unexpected error when mocking the database.")
	}

	mock.ExpectExec("UPDATE questions").
		WithArgs(1, 23).WillReturnResult(sqlmock.NewResult(23, 1))

	qStore := NewQuestionStore(db)
	err = qStore.UpdateLikeTotal(23, 1)
	assert.NoError(t, err)
}

func TestPostgresQuestionStore_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	mQuestion := domain.Question{
		Question: "How you doing?", QuestionerName: "Mathew", AssociatedRoom: "room1",
	}

	insertQuery := `INSERT INTO questions (question, questioner_name, fk_room_code) 
						VALUES ($1, $2, $3) 
						RETURNING question_id`
	mock.ExpectQuery(insertQuery).
		WithArgs(mQuestion.Question, mQuestion.QuestionerName, mQuestion.AssociatedRoom).
		WillReturnRows(sqlmock.NewRows([]string{"question_id"}).AddRow(1))

	qStore := NewQuestionStore(db)
	err = qStore.Create(&mQuestion)
	assert.NoError(t, err)
	assert.Equal(t, 1, mQuestion.QuestionId)
}

func TestPostgresQuestionStore_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	mock.ExpectExec("DELETE FROM questions").WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	qStore := NewQuestionStore(db)
	err = qStore.Delete(1)
	assert.NoError(t, err)
}
