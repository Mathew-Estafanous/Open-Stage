package domain

type Question struct {
	QuestionId     int    `json:"question_id"`
	Question       string `json:"question"`
	QuestionerName string `json:"questioner_name"`
	TotalLikes     int    `json:"total_likes"`
	AssociatedRoom string `json:"associated_room"`
}

// QuestionStore is an interface that describes the given contract
// that must be met for each question repo.
type QuestionStore interface {
	GetById(id int) (Question, error)
	GetAllInRoom(roomCode string) ([]Question, error)
	UpdateLikeTotal(id int, total int) error
	Delete(id int) error
	Create(q *Question) error
}

type QuestionService interface {
	FindWithId(id int) (Question, error)
	FindAllInRoom(code string) ([]Question, error)
	ChangeTotalLikes(id int, total int) error
	Create(q *Question) error
	Delete(id int) error
}
