package domain

type Code int

const (
	Internal Code = iota
	NotFound
	Conflict
	BadInput
	Unauthorized
)

func (c Code) Error() string {
	switch c {
	case BadInput:
		return "Bad Input"
	case NotFound:
		return "Not Found"
	case Conflict:
		return "Conflict"
	case Unauthorized:
		return "Unauthorized"
	default:
		return "Internal Error"
	}
}
