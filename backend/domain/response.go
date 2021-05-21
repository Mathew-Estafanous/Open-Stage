package domain

type Code int

const (
	Internal Code = iota
	NotFound
	Conflict
	BadInput
)

func (c Code) Error() string {
	switch c {
	case BadInput:
		return "Bad Input"
	case NotFound:
		return "Not Found"
	case Conflict:
		return "Conflict"
	default:
		return "Internal Error"
	}
}
