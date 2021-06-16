package domain

// Room models every group Q&A session/room.
//
// Every room contains one main host and a unique room code which
// then contains questions that pertain to the room and whatever
// topic is in discussion.
//
// swagger:model room
type Room struct {
	// Unique ID for the room.
	//
	// example: gopherCon
	RoomCode string `json:"room_code"`

	// The host of the room's name.
	//
	// required: true
	// example: Mathew
	Host string `json:"host"`

	// The account that owns the room.
	//
	// example: 3434
	AccId int `json:"account_id"`
}

// RoomStore is an interface that describes the given contract
// that must be met for each room repository.
type RoomStore interface {
	Delete(roomCode string) error
	GetByRoomCode(code string) (Room, error)
	Create(room *Room) error
	FindAllRooms(accId int) ([]Room, error)
}

type RoomService interface {
	DeleteRoom(code string, accId int) error
	CreateRoom(room *Room) error
	FindRoom(roomCode string) (Room, error)
	AllRoomsWithId(accId int) ([]Room, error)
}
