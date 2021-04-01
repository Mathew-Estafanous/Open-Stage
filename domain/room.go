package domain

type Room struct {
	RoomId   int `json:"room_id"`
	RoomCode string `json:"room_code"`
	Host     string `json:"host"`
}

// RoomStore is an interface that describes the given contract
// that must be met for each room repository.
type RoomStore interface {
	GetByRoomCode(code string) (Room, error)
	Create(room *Room) error
}
