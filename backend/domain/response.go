package domain

type ErrType int

const (
	BadInput ErrType = iota
	NotFound
	Conflict
	Internal
)

type ApiError struct {
	Msg string
	Typ ErrType
}

func (e ApiError) Error() string {
	return e.Msg
}
