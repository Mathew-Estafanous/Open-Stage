package domain

type Room struct {
	RoomId   string `json:"room_id"`
	RoomCode string `json:"room_code"`
	Host     string `json:"host"`
}

// RoomRepository is an interface that describes the given contract
// that must be met for each room repository.
type RoomRepository interface {
	GetByCode(code string) (Room, error)
	Create(room *Room) error
}
