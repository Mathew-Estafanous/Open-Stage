package domain

type Question struct {
	QuestionId     int    `json:"question_id"`
	Question       string `json:"question"`
	QuestionerName string `json:"questioner_name"`
	TotalLikes     string `json:"total_likes"`
	AssociatedRoom string `json:"associated_room"`
}

// QuestionRepository is an interface that describes the given contract
// that must be met for each question repo.
type QuestionRepository interface {
	GetById(id int) (Question, error)
	Create(q *Question) (id int, err error)
}
