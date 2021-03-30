package domain


type Room struct {
	RoomId string `json:"room_id"`
	Host   string `json:"host"`
}

// RoomRepository is an interface that describes the given contract
// that must be met for each room repository.
type RoomRepository interface {
	GetById(id int) (Room, error)
	Create(room *Room) (string, error)
}