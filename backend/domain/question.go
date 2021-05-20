package domain

// Question represents every question asked within a room.
//
// Every question is given an associated room that the question has been
// asked.
//
// swagger:model question
type Question struct {
	// The id for each question
	//
	// example: 3452
	QuestionId int `json:"question_id"`

	// The question that was asked.
	//
	// example: What is 2 + 2?
	Question string `json:"question"`

	// Name of the questioner.
	//
	// example: Anonymous
	QuestionerName string `json:"questioner_name"`

	// The total # of likes for that question.
	//
	// min: 0
	// example: 2
	TotalLikes int `json:"total_likes"`

	// The room that the question was asked in.
	//
	// example: conference20
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
	ChangeTotalLikes(id int, change int) (Question, error)
	Create(q *Question) error
	Delete(id int) error
}
