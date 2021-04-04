package database

import (
	"github.com/Mathew-Estafanous/Open-Stage/domain"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestMySQLQuestionStore_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	mQuestion := domain.Question{
		QuestionId: 2, Question: "How you doing?", QuestionerName: "Mathew", TotalLikes: 0, AssociatedRoom: "room1",
	}

	row := sqlmock.NewRows([]string{"question_id", "question", "questioner_name", "total_likes", "fk_room_id"}).
		AddRow(mQuestion.QuestionId, mQuestion.Question, mQuestion.QuestionerName,
			mQuestion.TotalLikes, mQuestion.AssociatedRoom)

	questionId := 2
	query := `SELECT question_id, question, questioner_name, total_likes, fk_room_id
				FROM questions WHERE question_id = ?`
	mock.ExpectQuery(query).WithArgs(questionId).WillReturnRows(row)

	qStore := NewMySQLQuestionStore(db)
	q, err := qStore.GetById(2)

	assert.NoError(t, err)
	assert.EqualValues(t, mQuestion, q)
}

func TestMySQLQuestionStore_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("There was an unexpected error when mocking the database.")
	}

	mQuestion := domain.Question{
		Question: "How you doing?", QuestionerName: "Mathew", AssociatedRoom: "room1",
	}

	mock.ExpectExec("INSERT INTO questions").
		WithArgs(&mQuestion.Question, &mQuestion.QuestionerName, &mQuestion.AssociatedRoom).
		WillReturnResult(sqlmock.NewResult(1, 1))

	qStore := NewMySQLQuestionStore(db)
	err = qStore.Create(&mQuestion)
	assert.NoError(t, err)
	assert.Equal(t, 1, mQuestion.QuestionId)
}
